// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package multiraft

import (
	"time"

	"golang.org/x/net/context"

	"github.com/journeymidnight/nentropy/helper"
	"github.com/journeymidnight/nentropy/multiraft/multiraftbase"
	"github.com/journeymidnight/nentropy/util/stop"
	"github.com/journeymidnight/nentropy/util/syncutil"
	"github.com/journeymidnight/nentropy/util/timeutil"
)

// A replicaQueue is a prioritized queue of replicas for which work is
// scheduled. For example, there's a GC queue for replicas which are due
// for garbage collection, a rebalance queue to move replicas from full
// or busy stores, a recovery queue for replicas of ranges with dead replicas,
// etc.
type replicaQueue interface {
	// Start launches a goroutine to process the contents of the queue.
	// The provided stopper is used to signal that the goroutine should exit.
	Start(*stop.Stopper)
	// MaybeAdd adds the replica to the queue if the replica meets
	// the queue's inclusion criteria and the queue is not already
	// too full, etc.
	MaybeAdd(*Replica)
	// MaybeRemove removes the replica from the queue if it is present.
	MaybeRemove(multiraftbase.GroupID)
}

// A replicaSet provides access to a sequence of replicas to consider
// for inclusion in replica queues. There are no requirements for the
// ordering of the iteration.
type replicaSet interface {
	// Visit calls the given function for every replica in the set btree
	// until the function returns false.
	Visit(func(*Replica) bool)
	// EstimatedCount returns the number of replicas estimated to remain
	// in the iteration. This value does not need to be exact.
	EstimatedCount() int
}

// A replicaScanner iterates over replicas at a measured pace in order to
// complete approximately one full scan per target interval in a large
// store (in small stores it may complete faster than the target
// interval).  Each replica is tested for inclusion in a sequence of
// prioritized replica queues.
type replicaScanner struct {
	helper.AmbientContext

	targetInterval time.Duration  // Target duration interval for scan loop
	maxIdleTime    time.Duration  // Max idle time for scan loop
	waitTimer      timeutil.Timer // Shared timer to avoid allocations.
	replicas       replicaSet     // Replicas to be scanned
	queues         []replicaQueue // Replica queues managed by this scanner
	removed        chan *Replica  // Replicas to remove from queues
	// Count of times and total duration through the scanning loop.
	mu struct {
		syncutil.Mutex
		scanCount        int64
		waitEnabledCount int64
		total            time.Duration
		// Some tests in this package disable scanning.
		disabled bool
	}
	// Used to notify processing loop if the disabled state changes.
	setDisabledCh chan struct{}
}

// newReplicaScanner creates a new replica scanner with the provided
// loop intervals, replica set, and replica queues.  If scanFn is not
// nil, after a complete loop that function will be called. If the
// targetInterval is 0, the scanner is disabled.
func newReplicaScanner(
	ambient helper.AmbientContext, targetInterval, maxIdleTime time.Duration, replicas replicaSet,
) *replicaScanner {
	if targetInterval < 0 {
		panic("scanner interval must be greater than or equal to zero")
	}
	rs := &replicaScanner{
		AmbientContext: ambient,
		targetInterval: targetInterval,
		maxIdleTime:    maxIdleTime,
		replicas:       replicas,
		removed:        make(chan *Replica, 10),
		setDisabledCh:  make(chan struct{}, 1),
	}
	if targetInterval == 0 {
		rs.SetDisabled(true)
	}
	return rs
}

// AddQueues adds a variable arg list of queues to the replica scanner.
// This method may only be called before Start().
func (rs *replicaScanner) AddQueues(queues ...replicaQueue) {
	rs.queues = append(rs.queues, queues...)
}

// Start spins up the scanning loop.
func (rs *replicaScanner) Start(stopper *stop.Stopper) {
	for _, queue := range rs.queues {
		queue.Start(stopper)
	}
	rs.scanLoop(stopper)
}

// scanCount returns the number of times the scanner has cycled through
// all replicas.
func (rs *replicaScanner) scanCount() int64 {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.mu.scanCount
}

// waitEnabledCount returns the number of times the scanner went in the mode of
// waiting to be reenabled.
func (rs *replicaScanner) waitEnabledCount() int64 {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.mu.waitEnabledCount
}

// SetDisabled turns replica scanning off or on as directed. Note that while
// disabled, removals are still processed.
func (rs *replicaScanner) SetDisabled(disabled bool) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.mu.disabled = disabled
	// The select prevents blocking on the channel.
	select {
	case rs.setDisabledCh <- struct{}{}:
	default:
	}
}

func (rs *replicaScanner) GetDisabled() bool {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.mu.disabled
}

// avgScan returns the average scan time of each scan cycle. Used in unittests.
func (rs *replicaScanner) avgScan() time.Duration {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	if rs.mu.scanCount == 0 {
		return 0
	}
	return time.Duration(rs.mu.total.Nanoseconds() / rs.mu.scanCount)
}

// RemoveReplica removes a replica from any replica queues the scanner may
// have placed it in. This method should be called by the Store
// when a replica is removed (e.g. rebalanced or merged).
func (rs *replicaScanner) RemoveReplica(repl *Replica) {
	rs.removed <- repl
}

// paceInterval returns a duration between iterations to allow us to pace
// the scan.
func (rs *replicaScanner) paceInterval(start, now time.Time) time.Duration {
	elapsed := now.Sub(start)
	remainingNanos := rs.targetInterval.Nanoseconds() - elapsed.Nanoseconds()
	if remainingNanos < 0 {
		remainingNanos = 0
	}
	count := rs.replicas.EstimatedCount()
	if count < 1 {
		count = 1
	}
	interval := time.Duration(remainingNanos / int64(count))
	if rs.maxIdleTime > 0 && interval > rs.maxIdleTime {
		interval = rs.maxIdleTime
	}
	return interval
}

// waitAndProcess waits for the pace interval and processes the replica
// if repl is not nil. The method returns true when the scanner needs
// to be stopped. The method also removes a replica from queues when it
// is signaled via the removed channel.
func (rs *replicaScanner) waitAndProcess(
	ctx context.Context, start time.Time, stopper *stop.Stopper, repl *Replica,
) bool {
	waitInterval := rs.paceInterval(start, timeutil.Now())
	rs.waitTimer.Reset(waitInterval)

	helper.Printf(5, "wait timer interval set to %s", waitInterval)
	for {
		select {
		case <-rs.waitTimer.C:
			helper.Printf(5, "wait timer fired")
			rs.waitTimer.Read = true
			if repl == nil {
				return false
			}

			helper.Printf(5, "replica scanner processing %s", repl)
			for _, q := range rs.queues {
				q.MaybeAdd(repl)
			}
			return false

		case repl := <-rs.removed:
			rs.removeReplica(repl)

		case <-stopper.ShouldStop():
			return true
		}
	}
}

func (rs *replicaScanner) removeReplica(repl *Replica) {
	// Remove replica from all queues as applicable. Note that we still
	// process removals while disabled.
	groupID := repl.GroupID
	for _, q := range rs.queues {
		q.MaybeRemove(groupID)
	}

	helper.Printf(5, "removed replica %s", repl)

}

// scanLoop loops endlessly, scanning through replicas available via
// the replica set, or until the scanner is stopped. The iteration
// is paced to complete a full scan in approximately the scan interval.
func (rs *replicaScanner) scanLoop(stopper *stop.Stopper) {
	ctx := rs.AnnotateCtx(context.Background())
	stopper.RunWorker(ctx, func(ctx context.Context) {
		start := timeutil.Now()

		// waitTimer is reset in each call to waitAndProcess.
		defer rs.waitTimer.Stop()

		for {
			if rs.GetDisabled() {
				if done := rs.waitEnabled(stopper); done {
					return
				}
				continue
			}
			var shouldStop bool
			count := 0
			rs.replicas.Visit(func(repl *Replica) bool {
				count++
				shouldStop = rs.waitAndProcess(ctx, start, stopper, repl)
				return !shouldStop
			})
			if count == 0 {
				// No replicas processed, just wait.
				shouldStop = rs.waitAndProcess(ctx, start, stopper, nil)
			}

			shouldStop = shouldStop || nil != stopper.RunTask(
				ctx, "storage.replicaScanner: scan loop",
				func(ctx context.Context) {
					// Increment iteration count.
					rs.mu.Lock()
					defer rs.mu.Unlock()
					rs.mu.scanCount++
					rs.mu.total += timeutil.Since(start)

					helper.Printf(5, "reset replica scan iteration")

					// Reset iteration and start time.
					start = timeutil.Now()
				})
			if shouldStop {
				return
			}
		}
	})
}

// waitEnabled loops, removing replicas from the scanner's queues,
// until scanning is enabled or the stopper signals shutdown,
func (rs *replicaScanner) waitEnabled(stopper *stop.Stopper) bool {
	rs.mu.Lock()
	rs.mu.waitEnabledCount++
	rs.mu.Unlock()
	for {
		if !rs.GetDisabled() {
			return false
		}
		select {
		case <-rs.setDisabledCh:
			continue

		case repl := <-rs.removed:
			rs.removeReplica(repl)

		case <-stopper.ShouldStop():
			return true
		}
	}
}
