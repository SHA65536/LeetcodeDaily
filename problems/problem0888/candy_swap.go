package problem0888

/*
Alice and Bob have a different total number of candies.
You are given two integer arrays aliceSizes and bobSizes where
aliceSizes[i] is the number of candies of the ith box of candy that Alice has and bobSizes[j] is the number of candies of the jth box of candy that Bob has.
Since they are friends, they would like to exchange one candy box each so that after the exchange, they both have the same total amount of candy.
The total amount of candy a person has is the sum of the number of candies in each box they have.
Return an integer array answer where answer[0] is the number of candies in the box that Alice must exchange,
and answer[1] is the number of candies in the box that Bob must exchange. If there are multiple answers, you may return any one of them.
It is guaranteed that at least one answer exists.
*/

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var aliceSum, bobSum int
	var bobMap = map[int]bool{}
	for i := range aliceSizes {
		aliceSum += aliceSizes[i]
	}
	for i := range bobSizes {
		bobSum += bobSizes[i]
		bobMap[bobSizes[i]] = true
	}
	for _, num := range aliceSizes {
		wanted := bobSum + (2 * num) - aliceSum
		if wanted%2 == 0 && bobMap[wanted/2] {
			return []int{num, wanted / 2}
		}
	}
	return nil
}

// Solve for x
// aliceSum-num+x = bobSum+num-x
// 2x = bobSum+2x-aliceSum
// x = (bobSum+2x-aliceSum)/2
