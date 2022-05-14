package problem0743

/*
https://leetcode.com/problems/network-delay-time/

You are given a network of n nodes, labeled from 1 to n.
You are also given 'times', a list of travel times as directed edges times[i] = (ui, vi, wi),
where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.
We will send a signal from a given node k.
Return the time it takes for all the n nodes to receive the signal.
If it is impossible for all the n nodes to receive the signal, return -1.
*/

const (
	src = iota
	dst
	dur
)

func networkDelayTime(times [][]int, n int, k int) int {
	var timesMap = map[int][][]int{}
	var visited = map[int]bool{}
	for _, rt := range times {
		if _, ok := timesMap[rt[src]]; ok {
			timesMap[rt[src]] = append(timesMap[rt[src]], []int{rt[src], rt[dst], rt[dur]})
		} else {
			timesMap[rt[src]] = [][]int{{rt[src], rt[dst], rt[dur]}}
		}
	}
	maxtime, values := getTimes(timesMap, visited, k)
	if float64(values) == (float64(n)/2)*float64(1+n) {
		return maxtime
	}
	return -1
}

func getTimes(timesMap map[int][][]int, visited map[int]bool, cur int) (time int, vals int) {
	vals = cur
	visited[cur] = true
	if routes, ok := timesMap[cur]; ok {
		for _, rt := range routes {
			if _, ok := visited[rt[dst]]; ok {
				continue
			}
			nextTime, val := getTimes(timesMap, visited, rt[dst])
			nextTime += rt[dur]
			vals += val
			if nextTime > time {
				time = nextTime
			}
		}
	}
	return
}
