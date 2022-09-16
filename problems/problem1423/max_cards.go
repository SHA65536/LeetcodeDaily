package problem1423

/*
https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards

There are several cards arranged in a row, and each card has an associated number of points.
The points are given in the integer array cardPoints.
In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
Your score is the sum of the points of the cards you have taken.
Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
*/

func maxScore(cardPoints []int, k int) int {
	// We are using a sliding window technique to find the
	// minimum sum group that lives inside cardPoints
	var subLength = len(cardPoints) - k
	var minSubsum, curSubsum, totalSum int
	// Getting sum of full group
	for i := 0; i < subLength; i++ {
		curSubsum += cardPoints[i]
		totalSum += cardPoints[i]
	}
	minSubsum = curSubsum
	// Moving the group slowly
	for i := subLength; i < len(cardPoints); i++ {
		curSubsum += cardPoints[i]
		totalSum += cardPoints[i]
		curSubsum -= cardPoints[i-subLength]
		if curSubsum < minSubsum {
			minSubsum = curSubsum
		}
	}
	// Subtracting the sum of min subgroup to get the sum of the outter cards
	return totalSum - minSubsum
}
