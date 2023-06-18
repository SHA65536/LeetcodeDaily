package problem1333

import "sort"

/*
Given the array restaurants where  restaurants[i] = [idi, ratingi, veganFriendlyi, pricei, distancei].
You have to filter the restaurants using three filters.
The veganFriendly filter will be either true(meaning you should only include restaurants with veganFriendlyi set to true)
or false (meaning you can include any restaurant).
In addition, you have the filters maxPrice and maxDistance which are the maximum value for price and distance of restaurants you should consider respectively.

Return the array of restaurant IDs after filtering, ordered by rating from highest to lowest.
For restaurants with the same rating, order them by id from highest to lowest.
For simplicity veganFriendly i and veganFriendly take value 1 when it is true, and 0 when it is false.
*/

func filterRestaurants(input [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	// res[i] will contain the index of the restaurant in input
	var res = make([]int, 0, len(input))
	for i, r := range input {
		// Filter by vegan
		if veganFriendly == 1 && r[2] != 1 {
			continue
		}
		// Filter by price and distance
		if r[3] > maxPrice || r[4] > maxDistance {
			continue
		}
		res = append(res, i)
	}
	// Sort by order with secondary by id
	sort.Slice(res, func(i, j int) bool {
		if input[res[i]][1] == input[res[j]][1] {
			return input[res[i]][0] > input[res[j]][0]
		}
		return input[res[i]][1] > input[res[j]][1]
	})
	// Replace index of restaurant with id
	for i := range res {
		res[i] = input[res[i]][0]
	}
	return res
}
