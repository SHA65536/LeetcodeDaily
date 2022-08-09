package problem0823

import "sort"

/*
Given an array of unique integers, arr, where each integer arr[i] is strictly greater than 1.
We make a binary tree using these integers, and each number may be used for any number of times.
Each non-leaf node's value should be equal to the product of the values of its children.
Return the number of binary trees we can make. The answer may be too large so return the answer modulo 109 + 7.
*/

const max = 1000000007

func numFactoredBinaryTrees(arr []int) int {
	var res int
	var cache = map[int]int{}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		cache[arr[i]] = 1
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				cache[arr[i]] = (cache[arr[i]] + cache[arr[j]]*cache[arr[i]/arr[j]]) % max
			}
		}
		res = (res + cache[arr[i]]) % max
	}
	return res
}
