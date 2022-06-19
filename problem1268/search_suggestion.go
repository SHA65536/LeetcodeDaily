package problem1268

import (
	"sort"
	"strings"
)

/*
https://leetcode.com/problems/search-suggestions-system/

You are given an array of strings products and a string searchWord.
Design a system that suggests at most three product names from products after each character of searchWord is typed.
Suggested products should have common prefix with searchWord.
If there are more than three products with a common prefix return the three lexicographically minimums products.
Return a list of lists of the suggested products after each character of searchWord is typed.
*/

func suggestedProducts(products []string, searchWord string) [][]string {
	var results = make([][]string, 0)
	var start int
	sort.Strings(products)
	for idx := range searchWord {
		results = append(results, []string{})
		if start == -1 {
			continue
		}
		prefix := searchWord[:idx+1]
		newStart := searchProducts(products, prefix, start)
		if newStart == -1 {
			continue
		}
		start = newStart
		for ; newStart < len(products) && newStart < start+3; newStart++ {
			if !strings.HasPrefix(products[newStart], prefix) {
				break
			}
			results[idx] = append(results[idx], products[newStart])
		}
	}
	return results
}

func searchProducts(products []string, prefix string, start int) int {
	var end, idx = len(products) - 1, -1
	for end >= start {
		mid := ((end - start) / 2) + start
		if strings.HasPrefix(products[mid], prefix) {
			if mid == 0 || !strings.HasPrefix(products[mid-1], prefix) {
				idx = mid
				break
			}
		}
		if prefix <= products[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return idx
}
