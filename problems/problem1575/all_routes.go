package problem1575

/*
You are given an array of distinct positive integers locations where locations[i] represents the position of city i.
You are also given integers start, finish and fuel representing the starting city, ending city, and the initial amount of fuel you have, respectively.
At each step, if you are at city i, you can pick any city j such that j != i and 0 <= j < locations.length and move to city j.
Moving from city i to city j reduces the amount of fuel you have by |locations[i] - locations[j]|.
Please notice that |x| denotes the absolute value of x.
Notice that fuel cannot become negative at any point in time, and that you are allowed to visit any city more than once (including start and finish).
Return the count of all possible routes from start to finish. Since the answer may be too large, return it modulo 109 + 7.
*/

const mod int = 1e9 + 7

func countRoutes(loc []int, src, dst, fuel int) int {
	// cache[i][f] represents the number of ways to get from
	// city i to the destination with f fuel left
	var cache = make([][]int, len(loc))
	// fill all with -1, so we know which we already calculated
	for i := range cache {
		cache[i] = make([]int, fuel+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return solveRoutes(cache, loc, src, dst, fuel)
}

func solveRoutes(cache [][]int, loc []int, src, dst, fuel int) int {
	var res int
	// Search in cache
	if cache[src][fuel] != -1 {
		return cache[src][fuel]
	}
	// Got to target
	if src == dst {
		res++
	}
	for nxt := range loc {
		// Can't stay in same town
		if nxt == src {
			continue
		}
		// Only go when you have enough fuel
		remaining := fuel - abs(loc[src]-loc[nxt])
		if remaining >= 0 {
			res = (res + solveRoutes(cache, loc, nxt, dst, remaining)) % mod
		}
	}
	cache[src][fuel] = res
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
