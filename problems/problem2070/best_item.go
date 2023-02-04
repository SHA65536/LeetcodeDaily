package problem2070

import "sort"

/*
You are given a 2D integer array items where items[i] = [pricei, beautyi] denotes the price and beauty of an item respectively.
You are also given a 0-indexed integer array queries.
For each queries[j], you want to determine the maximum beauty of an item whose price is less than or equal to queries[j].
If no such item exists, then the answer to this query is 0.
Return an array answer of the same length as queries where answer[j] is the answer to the jth query.
*/

func maximumBeauty(items [][]int, queries []int) []int {
	var max int
	var res = make([]int, len(queries))
	var priceIndex, valIndex []int
	// Sorting by increasing price, and decreasing value
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	// Removing all items with subpar price to value
	for i := range items {
		if items[i][1] > max {
			max = items[i][1]
			priceIndex = append(priceIndex, items[i][0])
			valIndex = append(valIndex, items[i][1])
		}
	}
	// Looking up queries with binary search
	for q := range queries {
		idx := sort.SearchInts(priceIndex, queries[q])
		if idx < len(priceIndex) && priceIndex[idx] == queries[q] {
			// Found exactly
			res[q] = valIndex[idx]
		} else if idx-1 >= 0 {
			// Not found, looking one cheaper
			res[q] = valIndex[idx-1]
		}
	}
	return res
}
