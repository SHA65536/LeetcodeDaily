package problem1615

/*
There is an infrastructure of n cities with some number of roads connecting these cities.
Each roads[i] = [ai, bi] indicates that there is a bidirectional road between cities ai and bi.
The network rank of two different cities is defined as the total number of directly connected roads to either city.
If a road is directly connected to both cities, it is only counted once.
The maximal network rank of the infrastructure is the maximum network rank of all pairs of different cities.
Given the integer n and the array roads, return the maximal network rank of the entire infrastructure.
*/

func maximalNetworkRank(n int, roads [][]int) int {
	var res int
	// Roads connected to each city
	var roadCnt = make([]int, n)
	// Roads shared by two cities
	var doubles = make(map[[2]int]int, len(roads))
	for _, rd := range roads {
		roadCnt[rd[0]]++
		roadCnt[rd[1]]++
		// Left is always lower
		doubles[[2]int{min(rd[0], rd[1]), max(rd[0], rd[1])}] = 1
	}
	// Try every city combo
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// All connected roads, minus shared roads
			cur := roadCnt[i] + roadCnt[j] - doubles[[2]int{i, j}]
			res = max(res, cur)
		}
	}
	return res
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
