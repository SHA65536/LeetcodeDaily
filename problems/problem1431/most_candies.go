package problem1431

/*
There are n kids with candies.
You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has,
and an integer extraCandies, denoting the number of extra candies that you have.
Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies,
they will have the greatest number of candies among all the kids, or false otherwise.
Note that multiple kids can have the greatest number of candies.
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res = make([]bool, len(candies))
	var max int

	for i := range candies {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := range candies {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}
