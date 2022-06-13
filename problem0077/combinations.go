package problem0077

/*
https://leetcode.com/problems/combinations/
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].

You may return the answer in any order.
*/

func combine(n int, k int) [][]int {
	var results = [][]int{}
	if k == 0 {
		return [][]int{[]int{}}
	}
	for i := k; i <= n; i++ {
		for _, combo := range combine(i-1, k-1) {
			combo = append(combo, i)
			results = append(results, combo)
		}
	}
	return results
}
