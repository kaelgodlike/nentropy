package store

import (
	"errors"
	"os"

	"github.com/dgraph-io/badger"
)

// Collection :wrapper of badger store
type Collection struct {
	Dir string
	kv  *badger.KV
}

// Iterator iterates a badger store
type Iterator struct {
	it *badger.Iterator
}

// WriteBatch provides Put/Delete methods like leveldb did
type WriteBatch struct {
	entries []*badger.Entry
}

// Onode holds the metadata of each object in osd store
type Onode struct {
	nid        uint64 //numeric id (locally unique)
	size       uint64 //object size
	stripeSize uint32 //size of each  stripe
}

// BadgerStore holds a number of Collection
// notice BadgerStore provides interfaces similar to leveldb and rocksdb, which is more popular
type BadgerStore struct {
	collections *[]Collection
}

// NewWriteBatch creates a new WriteBatch
func NewWriteBatch() *WriteBatch {
	return &WriteBatch{}
}

// Put add an key/value pair to this batch
func (wb *WriteBatch) Put(key, value []byte) {
	wb.entries = badger.EntriesSet(wb.entries, key, value)
}

// Delete add an key/value pair to this batch
func (wb *WriteBatch) Delete(key []byte) {
	wb.entries = badger.EntriesDelete(wb.entries, key)
}

// Length add an key/value pair to this batch
func (wb *WriteBatch) Length() int {
	return len(wb.entries)
}

// NewCollection return a store Collection, maybe create a new one
// if errorIfDirNotExists == true and dir not exists, raise error
func NewCollection(dir string, errorIfDirNotExists bool) (*Collection, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if errorIfDirNotExists == true {
			return nil, ErrDirNotExists
		}
		os.Mkdir(dir, 0755)
	}

	opt := badger.DefaultOptions
	opt.Dir = dir
	opt.ValueDir = dir
	kv, err := badger.NewKV(&opt)
	return &Collection{Dir: dir, kv: kv}, err
}

// Get returns value of an key, err on error
func (coll *Collection) Get(key []byte) (value []byte, err error) {
	var item badger.KVItem
	if geterr := coll.kv.Get(key, &item); geterr != nil {
		return value, errors.New("faild to get key")
	}
	var val []byte
	copyerr := item.Value(func(v []byte) {
		val = make([]byte, len(v))
		copy(val, v)
	})
	if copyerr != nil {
		return value, errors.New("faild to copy value")
	}

	return val, nil
}

// Remove remove a Collection's database, this is very dangerous
func (coll *Collection) Remove() {
	os.RemoveAll(coll.Dir)
}

// Put returns value of an key, err on error
func (coll *Collection) Put(key []byte, value []byte) (err error) {
	return coll.kv.Set(key, value, 0)
}

// Write writes entries to the badger storage
// fixme: maybe need better error handling
func (coll *Collection) Write(wb *WriteBatch) error {
	err := coll.kv.BatchSet(wb.entries)
	if err != nil {
		return err
	}

	for _, e := range wb.entries {
		if e.Error != nil {
			return e.Error
		}
	}

	return nil
}

// NewIterator create an Iterator
func (coll *Collection) NewIterator() *Iterator {
	return &Iterator{it: coll.kv.NewIterator(badger.IteratorOptions{})}
}

// Seek seek to @key
func (it *Iterator) Seek(key []byte) {
	it.it.Seek(key)
}

// Valid determines whether the iterator is currently valid
func (it *Iterator) Valid() bool {
	return it.it.Valid()
}

// Next move to next key
func (it *Iterator) Next() {
	it.it.Next()
}

// Close seek to @key
func (it *Iterator) Close() {
	it.it.Close()
}

// Rewind move iterator to beginning point
func (it *Iterator) Rewind() {
	it.it.Rewind()
}

// Key get the key points to Iterator
func (it *Iterator) Key() []byte {
	kvitem := it.it.Item()
	return kvitem.Key()
}

// Value get the value points to Iterator
func (it *Iterator) Value() (val []byte) {
	kvitem := it.it.Item()

	kvitem.Value(func(v []byte) {
		val = make([]byte, len(v))
		copy(val, v)
	})
	return
}

// Close closes an Collection
func (coll *Collection) Close() {
	coll.kv.Close()
}

// add an entry to an transaction
func (bs *BadgerStore) set(key, value []byte, entries []*badger.Entry) {
	entries = badger.EntriesSet(entries, key, value)
}

func stripeRead(offset uint64, value []byte) {

}
func stripeWrite(offset, length uint64, value []byte) {

}

// LoadCollections loads collections of this store
func LoadCollections() {

}