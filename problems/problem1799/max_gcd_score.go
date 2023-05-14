package problem1799

/*
You are given nums, an array of positive integers of size 2 * n.
You must perform n operations on this array.
In the ith operation (1-indexed), you will:
    Choose two elements, x and y.
    Receive a score of i * gcd(x, y).
    Remove x and y from nums.
Return the maximum score you can receive after performing n operations.
The function gcd(x, y) is the greatest common divisor of x and y.
*/

// maxScore utilises black magic to find the answer
func maxScore(nums []int) int {
	var dp = make([][]int, len(nums)/2+1)
	var dfs func(int, int) int
	for i := range dp {
		dp[i] = make([]int, 1<<len(nums))
	}
	dfs = func(i, mask int) int {
		if i > len(nums)/2 {
			return 0
		}
		if dp[i][mask] == 0 {
			for j := range nums {
				for k := j + 1; k < len(nums); k++ {
					new_mask := (1 << j) + (1 << k)
					if (mask & new_mask) == 0 {
						dp[i][mask] = max(dp[i][mask], i*gcd(nums[j], nums[k])+dfs(i+1, mask+new_mask))
					}
				}
			}
		}
		return dp[i][mask]
	}
	return dfs(1, 0)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
