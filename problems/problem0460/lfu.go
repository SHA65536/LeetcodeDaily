package problem0460

/*
Design and implement a data structure for a Least Frequently Used (LFU) cache.
Implement the LFUCache class:
LFUCache(int capacity) Initializes the object with the capacity of the data structure.
int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
void put(int key, int value) Update the value of the key if present, or inserts the key if not already present.
When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item.
For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
To determine the least frequently used key, a use counter is maintained for each key in the cache.
The key with the smallest use counter is the least frequently used key.
When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation).
The use counter for a key in the cache is incremented either a get or put operation is called on it.
The functions get and put must each run in O(1) average time complexity.
*/

import (
	"container/list"
)

type LFUCache struct {
	// Actual values in a map for O(n) access
	nodes map[int]*list.Element
	// Map of frequencies
	lists map[int]*list.List
	// Capacity of the cache
	cap int
	// The lowest frequency, for easy deletion
	min int
}

type Node struct {
	key, val, freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		make(map[int]*list.Element),
		make(map[int]*list.List),
		capacity, 0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	v, ok := lfu.nodes[key]
	if !ok {
		// If key not found, return -1
		return -1
	}
	n := v.Value.(*Node)
	// Remove the old node
	lfu.lists[n.freq].Remove(v)
	// Update frequency
	n.freq++
	// If frequency doesn't exists, create a new one
	if _, ok := lfu.lists[n.freq]; !ok {
		lfu.lists[n.freq] = list.New()
	}
	// Add the node at the back
	newlist := lfu.lists[n.freq]
	newnode := newlist.PushBack(n)
	lfu.nodes[key] = newnode
	// Update minimum
	if n.freq-1 == lfu.min && lfu.lists[n.freq-1].Len() == 0 {
		lfu.min++
	}
	return n.val
}

func (lfu *LFUCache) Put(key int, value int) {
	// If cap is 0, nothing we can put
	if lfu.cap == 0 {
		return
	}

	// If node exists, just update it
	if v, ok := lfu.nodes[key]; ok {
		n := v.Value.(*Node)
		n.val = value
		lfu.Get(key) // Calling get to update the frequency
		return
	}

	// If we're full, delete the first min frequency node
	if lfu.cap == len(lfu.nodes) {
		list := lfu.lists[lfu.min]
		frontNode := list.Front() // Front is the oldest
		delete(lfu.nodes, frontNode.Value.(*Node).key)
		list.Remove(frontNode)
	}

	// New freq is 1, guaranteed to be the min
	lfu.min = 1
	n := &Node{key, value, 1}

	// If this is a new frequency, create it
	if _, ok := lfu.lists[1]; !ok {
		lfu.lists[1] = list.New()
	}

	// Add the actual node
	list1 := lfu.lists[1]
	newnode := list1.PushBack(n)
	lfu.nodes[key] = newnode
}
