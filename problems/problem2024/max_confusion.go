package problem2024

/*
A teacher is writing a test with n true/false questions, with 'T' denoting true and 'F' denoting false.
He wants to confuse the students by maximizing the number of consecutive questions with the same answer (multiple trues or multiple falses in a row).
You are given a string answerKey, where answerKey[i] is the original answer to the ith question.
In addition, you are given an integer k, the maximum number of times you may perform the following operation:
    Change the answer key for any question to 'T' or 'F' (i.e., set answerKey[i] to 'T' or 'F').
Return the maximum number of consecutive 'T's or 'F's in the answer key after performing the operation at most k times.
*/

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(calcMax(answerKey, k+1, 'T'), calcMax(answerKey, k+1, 'F'))
}

func calcMax(answerKey string, k int, key byte) int {
	// idx is the current group index
	// sum is the sum of the k previous groups
	var res, idx, sum int

	// prev is the number of consecutive letters in the last k groups
	var prev = make([]int, k)

	for i := range answerKey {
		if answerKey[i] == key {
			// If it's the letter we're looking for, add to the current group
			prev[idx]++
			// Update the sum
			sum++
			// Update the max sum
			res = max(res, sum)
		} else {
			// If it's the wrong letter
			// Move to the next group
			idx = (idx + 1) % k
			// Remove the last group from the sum
			sum -= prev[idx]
			// Initialize the current group to 0
			prev[idx] = 0
		}
	}
	// Add the letters we could change to the sum
	return min(res+k-1, len(answerKey))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
