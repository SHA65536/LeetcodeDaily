package problem0985

/*
You are given an integer array nums and an array queries where queries[i] = [vali, indexi].
For each query i, first, apply nums[indexi] = nums[indexi] + vali, then print the sum of the even values of nums.
Return an integer array answer where answer[i] is the answer to the ith query.
*/

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	var sum int
	var sums = make([]int, len(queries))
	// Getting initial sum
	for i := range nums {
		if nums[i]%2 == 0 {
			sum += nums[i]
		}
	}
	for i := range queries {
		old := nums[queries[i][1]]
		new := old + queries[i][0]
		// If the old value is even, we want to
		// delete it from the sum
		if old%2 == 0 {
			sum -= old
		}
		// If the new value is even, we want to
		// add it to the sum
		if new%2 == 0 {
			sum += new
		}
		nums[queries[i][1]] = new
		sums[i] = sum
	}
	return sums
}
