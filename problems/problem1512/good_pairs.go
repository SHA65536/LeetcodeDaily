package problem1512

/*
Given an array of integers nums, return the number of good pairs.
A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/

func numIdenticalPairs(nums []int) int {
	var res int
	var freq = map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for _, v := range freq {
		if v > 1 {
			res += sumNumbers(v - 1)
		}
	}
	return res
}

func numIdenticalPairsArr(nums []int) int {
	var res int
	var freq [101]int
	for _, v := range nums {
		freq[v]++
	}
	for _, v := range freq {
		if v > 1 {
			res += sumNumbers(v - 1)
		}
	}
	return res
}

func sumNumbers(n int) int {
	return (n * (n + 1)) / 2
}
