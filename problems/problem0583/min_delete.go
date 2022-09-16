package problem0583

/*
https://leetcode.com/problems/delete-operation-for-two-strings/

Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.

In one step, you can delete exactly one character in either string.
*/

func minDistance(word1 string, word2 string) int {
	var memoization = make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		memoization[i] = make([]int, len(word2)+1)
	}

	//Finding the longest common subsequence by plotting all the
	//paths we can take and summing the common elements
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				memoization[i][j] = memoization[i-1][j-1] + 1
			} else {
				down := memoization[i-1][j]
				right := memoization[i][j-1]
				if down > right {
					memoization[i][j] = down
				} else {
					memoization[i][j] = right
				}
			}
		}
	}
	//Subtracting the total length from twice the longest subsequence
	return len(word1) + len(word2) - (2 * memoization[len(word1)][len(word2)])
}
