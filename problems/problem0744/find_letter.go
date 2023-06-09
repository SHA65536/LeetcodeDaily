package problem0744

/*
You are given an array of characters letters that is sorted in non-decreasing order, and a character target.
There are at least two different characters in letters.
Return the smallest character in letters that is lexicographically greater than target.
If such a character does not exist, return the first character in letters.
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	var low, high int = 0, len(letters) - 1
	for low < high {
		var mid = (low + high) / 2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if letters[low] <= target {
		return letters[0]
	}
	return letters[low]
}
