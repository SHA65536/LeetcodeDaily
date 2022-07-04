package problem0135

/*
https://leetcode.com/problems/candy/

There are n children standing in a line.
Each child is assigned a rating value given in the integer array ratings.
You are giving candies to these children subjected to the following requirements:
Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.
*/

func candy(ratings []int) int {
	var total int
	var candies []int = make([]int, len(ratings))
	// Checking each child's left side
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// Checking each child's right side
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			// Making sure not to break the left side
			if candies[i]+1 > candies[i-1] {
				candies[i-1] = candies[i] + 1
			}
		}
		// Each child has to get atleast 1
		total += candies[i] + 1
	}
	return total + candies[0] + 1
}
