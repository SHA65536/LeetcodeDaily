package problem1470

/*
Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
Return the array in the form [x1,y1,x2,y2,...,xn,yn].
*/

func shuffle(nums []int, n int) []int {
	var res = make([]int, len(nums))
	for i := 0; i < n; i++ {
		res[i*2] = nums[i]
		res[i*2+1] = nums[i+n]
	}
	return res
}
