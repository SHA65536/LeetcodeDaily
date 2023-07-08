package problem2551

import "sort"

/*
You have k bags. You are given a 0-indexed integer array weights where weights[i] is the weight of the ith marble.
You are also given the integer k.
Divide the marbles into the k bags according to the following rules:
    No bag is empty.
    If the ith marble and jth marble are in a bag, then all marbles with an index between the ith and jth indices should also be in that same bag.
    If a bag consists of all the marbles with an index from i to j inclusively, then the cost of the bag is weights[i] + weights[j].
The score after distributing the marbles is the sum of the costs of all the k bags.
Return the difference between the maximum and minimum scores among marble distributions.
*/

func putMarbles(w []int, k int) int64 {
	var res int64
	var pSum = make([]int64, len(w)-1)
	for i := 0; i < len(w)-1; i++ {
		pSum[i] = int64(w[i] + w[i+1])
	}
	sort.Slice(pSum, func(i, j int) bool { return pSum[i] < pSum[j] })
	for i := 0; i < k-1; i++ {
		res += pSum[len(pSum)-1-i] - pSum[i]
	}
	return res
}
