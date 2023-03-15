package problem2562

/*
You are given a 0-indexed integer array nums.
The concatenation of two numbers is the number formed by concatenating their numerals.
    For example, the concatenation of 15, 49 is 1549.
The concatenation value of nums is initially equal to 0. Perform this operation until nums becomes empty:
    If there exists more than one number in nums,
		pick the first element and last element in nums respectively and add the value of their concatenation to the concatenation value of nums,
		then delete the first and last element from nums.
    If one element exists, add its value to the concatenation value of nums, then delete it.
Return the concatenation value of the nums.
*/

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	for i := 0; i < len(nums)/2; i++ {
		res += concat(int64(nums[i]), int64(nums[len(nums)-i-1]))
	}
	if len(nums)%2 == 1 {
		res += int64(nums[len(nums)/2])
	}
	return res
}

func concat(a, b int64) int64 {
	var p int64
	for p = 1; p <= b; p *= 10 {
	}
	return (a * p) + b
}
