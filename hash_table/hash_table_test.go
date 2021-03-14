package hash_table

import (
	"fmt"
	"testing"
)

func populateHashTable(count int, start int) *ValueHashTable {
	dict := ValueHashTable{}
	for i := start; i < (start + count); i++ {
		dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}

func TestPut(t *testing.T) {
	dict := populateHashTable(3, 0)
	size := dict.Size()
	if size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	// should not add a new one, just change the existing one
	dict.Put("key1", "value1")
	size = dict.Size()
	if size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	dict.Put("key4", "value4")
	size = dict.Size()
	if size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}

func TestRemove(t *testing.T) {
	dict := populateHashTable(3, 0)
	dict.Remove("key2")
	size := dict.Size()
	if size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestGet(t *testing.T) {
	dict := populateHashTable(3, 0)
	value := dict.Get("key1")
	if value != "value1" {
		t.Errorf("wrong value, expect value to be value1")
	}
}
