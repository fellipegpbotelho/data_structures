package hash_table

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Key is the key of the dictionary
type Key generic.Type

// Value is the content of the dictionary
type Value generic.Type

// ValueHashTable is the set of items
type ValueHashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

// Uses the Horner's method to generate a hash of a string with O(n) complexity
func hash(key Key) int {
	keyFormatted := fmt.Sprintf("%s", key)
	h := 0
	for i := 0; i < len(keyFormatted); i++ {
		h = 31*h + int(keyFormatted[i])
	}
	return h
}

// Put item with value v and key k into the hash table
func (hashTable *ValueHashTable) Put(key Key, value Value) {
	hashTable.lock.Lock()
	defer hashTable.lock.Unlock()

	index := hash(key)
	if hashTable.items == nil {
		hashTable.items = make(map[int]Value)
	}
	hashTable.items[index] = value
}

// Remove item with key k from hash table
func (hashTable *ValueHashTable) Remove(key Key) {
	hashTable.lock.Lock()
	defer hashTable.lock.Unlock()

	index := hash(key)
	delete(hashTable.items, index)
}

// Get item with key k from the hash table
func (hashTable *ValueHashTable) Get(key Key) Value {
	hashTable.lock.RLock()
	defer hashTable.lock.RUnlock()

	index := hash(key)
	return hashTable.items[index]
}

// Size returns the number of the hash table elements
func (hashTable *ValueHashTable) Size() int {
	hashTable.lock.RLock()
	defer hashTable.lock.RUnlock()

	return len(hashTable.items)
}
