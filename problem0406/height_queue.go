package problem0406

import (
	"sort"
)

/*
https://leetcode.com/problems/queue-reconstruction-by-height

You are given an array of people, people, which are the attributes of some people in a queue (not necessarily in order).
Each people[i] = [hi, ki] represents the ith person of height hi with exactly ki other people in front who have a height greater than or equal to hi.
Reconstruct and return the queue that is represented by the input array people.
The returned queue should be formatted as an array queue,
where queue[j] = [hj, kj] is the attributes of the jth person in the queue (queue[0] is the person at the front of the queue).
*/

func reconstructQueue(people [][]int) [][]int {
	var results = make([][]int, 0)
	// Sorting people based on height then on number of people
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// Inserting people to the results based on height
	for _, person := range people {
		results = append(results, nil)
		copy(results[person[1]+1:], results[person[1]:])
		results[person[1]] = person
	}
	return results
}
