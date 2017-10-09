package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleSetGet(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	key := []byte("hello")
	value := []byte("world")
	err = coll.Put(key, value)
	require.Equal(t, err, nil)

	newvalue, err := coll.Get(key)
	require.Equal(t, newvalue, value)
	require.Equal(t, err, nil)
	coll.Close()
	coll.Remove()
}

func TestSimpleIterator(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	n := 100
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		err = coll.Put(k, k)
		require.Equal(t, err, nil)
	}

	itr := coll.NewIterator()
	i := 0
	for itr.Rewind(); itr.Valid(); itr.Next() {
		require.Equal(t, itr.Value(), []byte(fmt.Sprintf("%04d", i)))
		i++
	}
	coll.Remove()
}

func TestIteratorSeeking(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	n := 100
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		err = coll.Put(k, k)
		require.Equal(t, err, nil)
	}

	itr := coll.NewIterator()
	tobeseeked := []byte(fmt.Sprintf("%04d", 55))
	itr.Seek(tobeseeked)

	require.Equal(t, itr.Valid(), true)
	require.Equal(t, itr.Value(), tobeseeked)
	coll.Remove()
}

func TestWriteBatch_AllPuts(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	wb := NewWriteBatch()

	n := 100
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		wb.Put(k, k)
	}

	err = coll.Write(wb)
	require.Equal(t, err, nil)

	itr := coll.NewIterator()
	i := 0
	for itr.Rewind(); itr.Valid(); itr.Next() {
		require.Equal(t, itr.Value(), []byte(fmt.Sprintf("%04d", i)))
		i++
	}

	coll.Remove()
}
func TestWriteBatch_AllDeletes(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	n := 100
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		coll.Put(k, k)
	}

	wb := NewWriteBatch()
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		wb.Delete(k)
	}
	err = coll.Write(wb)
	require.Equal(t, err, nil)

	itr := coll.NewIterator()
	i := 0
	for itr.Rewind(); itr.Valid(); itr.Next() {
		i++
	}

	require.Equal(t, i, 0)
	coll.Remove()
}
func TestWriteBatch_DeleteAll(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	wb := NewWriteBatch()

	n := 100
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		wb.Put(k, k)
	}

	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		wb.Delete(k)
	}

	err = coll.Write(wb)
	require.Equal(t, err, nil)

	itr := coll.NewIterator()
	i := 0
	for itr.Rewind(); itr.Valid(); itr.Next() {
		require.Equal(t, itr.Value(), []byte(fmt.Sprintf("%04d", i)))
		i++
	}
	require.Equal(t, i, 0)

	coll.Remove()
}

func TestWriteBatch_DeleteHalf(t *testing.T) {
	coll, err := NewCollection("asdf", false)
	require.Equal(t, err, nil)

	wb := NewWriteBatch()

	n := 100
	for i := 0; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		wb.Put(k, k)
	}

	for i := 50; i < n; i++ {
		k := []byte(fmt.Sprintf("%04d", i))
		wb.Delete(k)
	}

	err = coll.Write(wb)
	require.Equal(t, err, nil)

	itr := coll.NewIterator()
	i := 0
	for itr.Rewind(); itr.Valid(); itr.Next() {
		require.Equal(t, itr.Value(), []byte(fmt.Sprintf("%04d", i)))
		i++
	}
	require.Equal(t, i, 50)

	coll.Remove()
}