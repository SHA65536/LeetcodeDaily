package problem1187

import (
	"math"
	"sort"
)

/*
Given two integer arrays arr1 and arr2, return the minimum number of operations (possibly zero) needed to make arr1 strictly increasing.
In one operation, you can choose two indices 0 <= i < arr1.length and 0 <= j < arr2.length and do the assignment arr1[i] = arr2[j].
If there is no way to make arr1 strictly increasing, return -1.
*/

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	var dp = map[int]int{-1: 0}
	for _, a1 := range arr1 {
		var temp = map[int]int{}
		for key, val := range dp {
			if a1 > key {
				last := math.MaxInt
				if v, ok := temp[a1]; ok {
					last = v
				}
				temp[a1] = min(last, val)
			}
			var idx = binarySearch(arr2, key)
			if idx < len(arr2) {
				last := math.MaxInt
				if v, ok := temp[arr2[idx]]; ok {
					last = v
				}
				temp[arr2[idx]] = min(last, val+1)
			}
		}
		dp = temp
	}
	var res = math.MaxInt
	if len(dp) == 0 {
		return -1
	}
	for _, val := range dp {
		res = min(res, val)
	}
	return res
}

func binarySearch(in []int, t int) int {
	var low, high = 0, len(in)
	for low < high {
		mid := low + (high-low)/2
		if in[mid] <= t {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
