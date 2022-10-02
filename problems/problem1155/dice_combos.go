package problem1155

/*
You have n dice and each die has k faces numbered from 1 to k.
Given three integers n, k, and target, return the number of possible ways
(out of the kn total ways) to roll the dice so the sum of the face-up numbers equals target.
Since the answer may be too large, return it modulo 109 + 7.
*/

const MOD = 1000000007

func numRollsToTarget(n int, k int, target int) int {
	// cache[i,j] represents the number of wayts to roll 'i' dice with sum 'j'
	var cache = map[[2]int]int{}
	if n*k < target {
		return 0
	}
	if n*k == target {
		return 1
	}
	return numRollsDp(cache, k, n, target)
}

func numRollsDp(cache map[[2]int]int, faces, dice, target int) int {
	var res int
	// check if we reached the end
	if dice == 0 && target == 0 {
		return 1
	} else if target <= 0 || (dice == 0 && target != 0) {
		return 0
	}
	// check in cache first
	if val, ok := cache[[2]int{dice, target}]; ok {
		return val
	}
	// for each face of the current dice, check subsequent combinations
	for i := 1; i <= faces; i++ {
		res += numRollsDp(cache, faces, dice-1, target-i)
	}
	// save result in cache
	cache[[2]int{dice, target}] = res % MOD
	return res % MOD
}
