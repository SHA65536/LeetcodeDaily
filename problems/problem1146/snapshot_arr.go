package problem1146

import "sort"

/*
Implement a SnapshotArray that supports the following interface:
SnapshotArray(int length) initializes an array-like data structure with the given length. Initially, each element equals 0.
void set(index, val) sets the element at the given index to be equal to val.
int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id
*/

type SnapVal struct {
	Id  int
	Val int
}

type SnapshotArray struct {
	// Snaps[i] represents all the values in index i at different snapshots
	Snaps [][]*SnapVal
	Size  int
	Cur   int
}

func Constructor(length int) SnapshotArray {
	snap := SnapshotArray{
		Snaps: make([][]*SnapVal, length),
		Size:  length,
		Cur:   0,
	}
	// Seeding the first snapshot with 0's
	for i := 0; i < length; i++ {
		snap.Snaps[i] = []*SnapVal{{0, 0}}
	}
	return snap
}

func (this *SnapshotArray) Set(index int, val int) {
	if this.Snaps[index][len(this.Snaps[index])-1].Id == this.Cur {
		// If snap id is the same as the last id (There was no snap), change the value
		this.Snaps[index][len(this.Snaps[index])-1].Val = val
	} else {
		// If snap id is different (There was a snap), add a new snap value
		this.Snaps[index] = append(this.Snaps[index], &SnapVal{this.Cur, val})
	}
}

func (this *SnapshotArray) Snap() int {
	this.Cur++
	return this.Cur - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// List of snap values at the given index
	indexSnap := this.Snaps[index]
	// Binary search showing the smallest index with bigger id
	latestIndex := sort.Search(len(indexSnap), func(i int) bool { return indexSnap[i].Id >= snap_id })
	if latestIndex >= len(indexSnap) || indexSnap[latestIndex].Id != snap_id {
		// Edge cases for when index is overflowing or not found
		latestIndex--
	}
	return this.Snaps[index][latestIndex].Val
}
