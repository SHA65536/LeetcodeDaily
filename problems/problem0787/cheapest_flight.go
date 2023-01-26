package problem0787

import (
	"math"
)

/*
There are n cities connected by some number of flights.
You are given an array flights where flights[i] = [from i, to i, price i] indicates that there is a flight from city fromi to city toi with cost price i.
You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops. If there is no such route, return -1.
*/

type Flight struct {
	Dst int
	Prc int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var res int = math.MaxInt
	var graph = map[int][]Flight{}
	var cur, nxt []Flight
	for _, f := range flights {
		graph[f[0]] = append(graph[f[0]], Flight{f[1], f[2]})
	}
	cur = []Flight{{src, 0}}
	for len(cur) > 0 && k >= 0 {
		k--
		nxt = []Flight{}
		for _, f := range cur {
			for _, d := range graph[f.Dst] {
				if d.Dst == dst {
					res = min(res, f.Prc+d.Prc)
				} else if f.Prc+d.Prc < res && f.Dst != d.Dst {
					nxt = append(nxt, Flight{d.Dst, f.Prc + d.Prc})
				}
			}
		}
		cur = nxt
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}

func findCheapestPriceOpt(n int, flights [][]int, src int, dst int, k int) int {
	var dp = make([]int, n)
	for i := range dp {
		dp[i] = 100000000
	}
	dp[src] = 0
	for z := 0; z <= k; z++ {
		temp := make([]int, n)
		copy(temp, dp)
		for _, f := range flights {
			temp[f[1]] = min(temp[f[1]], dp[f[0]]+f[2])
		}
		dp = temp
	}
	if dp[dst] == 100000000 {
		return -1
	}
	return dp[dst]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
