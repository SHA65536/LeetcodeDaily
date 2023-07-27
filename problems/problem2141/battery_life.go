package problem2141

import "sort"

/*
You have n computers. You are given the integer n and a 0-indexed integer array batteries
where the ith battery can run a computer for batteries[i] minutes.
You are interested in running all n computers simultaneously using the given batteries.
Initially, you can insert at most one battery into each computer.
After that and at any integer time moment, you can remove a battery from a computer and insert another battery any number of times.
The inserted battery can be a totally new battery or a battery from another computer.
You may assume that the removing and inserting processes take no time.
Note that the batteries cannot be recharged.
Return the maximum number of minutes you can run all the n computers simultaneously.
*/

func maxRunTime(n int, bats []int) int64 {
	sort.Ints(bats)
	var batSum = sum(bats)
	var k int
	for k = 0; int64(bats[len(bats)-1-k]) > batSum/int64(n-k); k++ {
		batSum -= int64(bats[len(bats)-1-k])
	}
	return batSum / int64(n-k)
}

func maxRunTimeBinSearch(n int, bats []int) int64 {
	var low, high int64 = 0, sum(bats) / int64(n)
	for low < high {
		mid := (high + low + 1) / 2
		if mid*int64(n) <= runTime(mid, bats) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func runTime(mid int64, bats []int) int64 {
	var res int64
	for _, b := range bats {
		res += min(int64(b), mid)
	}
	return res
}

func sum(bats []int) int64 {
	var res int64
	for _, b := range bats {
		res += int64(b)
	}
	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
