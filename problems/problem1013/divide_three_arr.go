package problem1013

/*
Given an array of integers arr, return true if we can partition the array into three non-empty parts with equal sums.
Formally, we can partition the array if we can find indexes i + 1 < j
with (arr[0] + arr[1] + ... + arr[i] == arr[i + 1] + arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1])
*/

func canThreePartsEqualSum(arr []int) bool {
	var sum, first, second, curSum int
	// Calculating overall sum
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	// If sum can't be divided into 3, reject
	if sum%3 != 0 {
		return false
	}
	sum /= 3
	// Finding first array
	for i := 0; i < len(arr); i++ {
		curSum += arr[i]
		if curSum == sum {
			first = i
			break
		}
	}
	// Finding second array
	curSum = 0
	for i := first + 1; i < len(arr); i++ {
		curSum += arr[i]
		if curSum == sum {
			second = i
			break
		}
	}
	// Finding last array
	curSum = 0
	for i := second + 1; i < len(arr); i++ {
		curSum += arr[i]
	}
	return curSum == sum && first < second && second < len(arr)-1
}
