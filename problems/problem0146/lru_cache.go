package problem0146

import "container/list"

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
Implement the LRUCache class:
    LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
    int get(int key) Return the value of the key if the key exists, otherwise return -1.
    void put(int key, int value) Update the value of the key if the key exists.
	Otherwise, add the key-value pair to the cache.
	If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.
*/

type Pair struct {
	Key, Val int
}

type LRUCache struct {
	// Doubly linked list of items in the cache
	// first item is most recently used
	// last item is least recently used
	Usage *list.List

	// Hashmap for accessing items at amortized O(1)
	Data map[int]*list.Element

	// Max size of the cache
	Cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Usage: list.New(),
		Data:  make(map[int]*list.Element, capacity),
		Cap:   capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if _, ok := lru.Data[key]; !ok {
		// Not Found
		return -1
	}

	// Move the item to the beginning of the list
	lru.Usage.MoveToFront(lru.Data[key])
	return lru.Usage.Front().Value.(Pair).Val
}

func (lru *LRUCache) Put(key int, value int) {
	// Found in data, Update value
	if val, ok := lru.Data[key]; ok {
		// Move to most used position
		lru.Usage.MoveToFront(val)
		// Update value
		val.Value = Pair{key, value}
		return
	}

	// Cache is full, discard least used element
	if lru.Usage.Len() == lru.Cap {
		// Remove least used
		last := lru.Usage.Back()
		lru.Usage.Remove(last)
		delete(lru.Data, last.Value.(Pair).Key)
	}

	// Add new key value to most used
	lru.Usage.PushFront(Pair{key, value})
	lru.Data[key] = lru.Usage.Front()
}
