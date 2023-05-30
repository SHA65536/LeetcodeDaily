package problem0705

/*
Design a HashSet without using any built-in hash table libraries.
Implement MyHashSet class:
    void add(key) Inserts the value key into the HashSet.
    bool contains(key) Returns whether the value key exists in the HashSet or not.
    void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.
*/

// Linked List to store the values of the hash
type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

// Search returns the node and the node before the value
func (list *LinkedNode) Search(val int) (res, prev *LinkedNode) {
	// Searching for value
	for res = list; res != nil; res = res.Next {
		// If found, return
		if res.Val == val {
			return res, prev
		}
		// Keep reference for last
		prev = res
	}

	// Return last if not found
	return res, prev
}

// Add adds a value to the linked list
func (list *LinkedNode) Add(val int) {
	var res, prev = list.Search(val)

	// If found, ignore
	if res != nil && res.Val == val {
		return
	}

	// Add to list
	prev.Next = &LinkedNode{Val: val}
}

// Remove removes a value from the linked list, returns the new head
func (list *LinkedNode) Remove(val int) *LinkedNode {
	var res, prev = list.Search(val)

	// If not found, ignore
	if res == nil {
		return list
	}

	if prev == nil && res == nil {
		// If it's the first and only item
		return nil
	} else if prev == nil && res != nil {
		// If it's the first item but not the only item
		return res.Next
	} else {
		// If it's not the first item, remove
		prev.Next = res.Next
	}
	return list
}

type MyHashSet struct {
	items []*LinkedNode
}

func Constructor() MyHashSet {
	// 200 is the size of the hash set, chosen arbitrarily, could be anything
	return MyHashSet{make([]*LinkedNode, 200)}
}

func (set *MyHashSet) Add(key int) {
	idx := key % len(set.items)
	if set.items[idx] == nil {
		set.items[idx] = &LinkedNode{Val: key}
	} else {
		set.items[idx].Add(key)
	}
}

func (set *MyHashSet) Remove(key int) {
	idx := key % len(set.items)
	if set.items[idx] != nil {
		set.items[idx] = set.items[idx].Remove(key)
	}
}

func (set *MyHashSet) Contains(key int) bool {
	idx := key % len(set.items)
	if set.items[idx] == nil {
		return false
	}
	res, _ := set.items[idx].Search(key)
	return res != nil
}
