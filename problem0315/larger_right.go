package problem0315

/*
You are given an integer array nums and you have to return a new counts array.
The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].
*/

func countSmallerNaive(nums []int) []int {
	var counts = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var count int
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		counts[i] = count
	}
	return counts
}

// We will use insertion sort with binary search to optimize
func countSmallerInsertion(nums []int) []int {
	var counts = make([]int, len(nums))
	var sorted = []int{}
	// Inserting each element
	for i := len(nums) - 1; i >= 0; i-- {
		// Binary search
		var start, end int = 0, len(sorted)
		for start < end {
			var mid int = start + (end-start)/2
			if sorted[mid] >= nums[i] {
				end = mid
			} else {
				start = mid + 1
			}
		}
		counts[i] = start
		if len(sorted) == start {
			sorted = append(sorted, nums[i])
		} else {
			sorted = append(sorted[:start+1], sorted[start:]...)
			sorted[start] = nums[i]
		}
	}
	return counts
}

/*
func countSmaller(nums []int) []int {
	var res = make([]int, len(nums))
	var idxs = make([]NumIdx, len(nums))
	for i := range nums {
		idxs[i].Idx = i
		idxs[i].Val = nums[i]
	}
	sort.Slice(idxs, func(i, j int) bool {
		return idxs[i].Val < idxs[j].Val
	})
	var prev, previdx int
	for i := len(idxs) - 1; i >= 0; i-- {
		if prev == idxs[i].Val {
			res[i] = previdx
		} else {
			res[i] = idxs[i].Idx
			previdx = idxs[i].Idx
		}
		prev = idxs[i].Val
	}
	return res
}

type NumIdx struct {
	Val, Idx int
}
*/
