package problem0967

/*
Return all non-negative integers of length n such that the absolute difference between every two consecutive digits is k.
Note that every number in the answer must not have leading zeros. For example, 01 has one leading zero and is invalid.
You may return the answer in any order.
*/

func numsSameConsecDiff(n int, k int) []int {
	var results = make([]int, 0)
	// Checking each starting digits excluding 0 as we're not allowed leading zeros
	for i := 1; i < 10; i++ {
		calcDigitDiff(&results, i, n-1, k)
	}
	return results
}

func calcDigitDiff(results *[]int, cur, left, diff int) {
	if left == 0 {
		// If we reached the desired length, add to results
		*results = append(*results, cur)
		return
	}
	if (cur%10)+diff <= 9 {
		// If adding the difference results in a valid number, keep checking
		calcDigitDiff(results, (cur*10)+(cur%10)+diff, left-1, diff)
	}
	// If the diffrence is 0, subtracting and adding will result in the same number
	if diff != 0 && (cur%10)-diff >= 0 {
		// If subtracting the difference results in a valid number, keep checking
		calcDigitDiff(results, (cur*10)+(cur%10)-diff, left-1, diff)
	}
}
