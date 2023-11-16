package problem1980

/*
Given an array of strings nums containing n unique binary strings each of length n,
return a binary string of length n that does not appear in nums.
If there are multiple answers, you may return any of them.
*/

func findDifferentBinaryString(nums []string) string {
	var res = make([]byte, len(nums))
	for i := range nums {
		if nums[i][i] == '1' {
			res[i] = '0'
		} else {
			res[i] = '1'
		}
	}
	return string(res)
}
