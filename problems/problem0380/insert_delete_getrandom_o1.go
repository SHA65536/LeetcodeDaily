package problem0380

import "math/rand"

/*
Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called).
Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.
*/

type RandomizedSet struct {
	ValOrder []int       // Keeps the values of the set
	ValMap   map[int]int // Map[n] contains the index of n in ValOrder
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		ValOrder: make([]int, 0),
		ValMap:   make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// If already exists, return false
	if _, ok := this.ValMap[val]; ok {
		return false
	}
	// Adding to the end of ValOrder
	this.ValOrder = append(this.ValOrder, val)
	// Updating the map accordingly
	this.ValMap[val] = len(this.ValOrder) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// If exists in set
	if idx, ok := this.ValMap[val]; ok {
		// Copy the last element to the location we want to remove
		this.ValOrder[idx] = this.ValOrder[len(this.ValOrder)-1]
		// Update the last element's index in map
		this.ValMap[this.ValOrder[idx]] = idx
		// Shrink the array because we moved the last element
		this.ValOrder = this.ValOrder[:len(this.ValOrder)-1]
		// Delete original element from map
		delete(this.ValMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	// Return random element at index 0 to length of ValOrder
	return this.ValOrder[rand.Intn(len(this.ValOrder))]
}
