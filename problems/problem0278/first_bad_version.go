package problem0278

/*
https://leetcode.com/problems/first-bad-version/

You are a product manager and currently leading a team to develop a new product.
Unfortunately, the latest version of your product fails the quality check.
Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad.
Implement a function to find the first bad version. You should minimize the number of calls to the API.
*/

func firstBadVersion(n int) int {
	var start, end = 0, n
	for end >= start {
		mid := ((end - start) / 2) + start
		if isBadVersion(mid) {
			if mid == 0 || !isBadVersion(mid-1) {
				return mid
			}
			end = mid - 1
		} else {
			if mid == n-1 || isBadVersion(mid+1) {
				return mid + 1
			}
			start = mid
		}
	}
	return 0
}

func isBadVersion(version int) bool {
	return true
}
