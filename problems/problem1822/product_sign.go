package problem1822

/*
There is a function signFunc(x) that returns:
    1 if x is positive.
    -1 if x is negative.
    0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.
Return signFunc(product).
*/

func arraySign(nums []int) int {
	var sign bool = true
	for _, n := range nums {
		if n == 0 {
			return 0
		} else if n < 0 {
			sign = !sign
		}
	}
	if sign {
		return 1
	}
	return -1
}
