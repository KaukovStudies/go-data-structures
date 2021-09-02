package hashtable

import "gitlab.com/Kaukov/go-data-structures/linked_list"

const TableSize int = 7

type HashTable struct {
	Data [TableSize]*linkedlist.LinkedList
}

func NewHashTable() *HashTable {
	ht := &HashTable{}

	for i := range ht.Data {
		ht.Data[i] = &linkedlist.LinkedList{}
	}

	return ht
}

func (ht *HashTable) Insert(key string) {
	index := ht.hash(key)

	if !ht.Search(key) {
		ht.Data[index].Append(&linkedlist.Node{Data: key})
	}
}

func (ht *HashTable) Search(key string) bool {
	index := ht.hash(key)

	return ht.Data[index].Search(key)
}

func (ht *HashTable) Delete(key string) {
	index := ht.hash(key)

	ht.Data[index].Delete(key)
}

func (ht HashTable) hash(key string) int {
	sum := 0

	for _, c := range key {
		sum += int(c)
	}

	return sum % TableSize
}
