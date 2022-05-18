package problem1192

/*
https://leetcode.com/problems/critical-connections-in-a-network/

There are n servers numbered from 0 to n - 1 connected by undirected server-to-server connections forming a network where connections[i] = [ai, bi] represents a connection between servers ai and bi.
Any server can reach other servers directly or indirectly through the network.
A critical connection is a connection that, if removed, will make some servers unable to reach some other server.
Return all critical connections in the network in any order.
*/

func criticalConnections(n int, connections [][]int) [][]int {
	var res = [][]int{}
	var nodes = map[int][]int{}
	for _, conn := range connections {
		nodes[conn[0]] = append(nodes[conn[0]], conn[1])
		nodes[conn[1]] = append(nodes[conn[1]], conn[0])
	}
	return res
}
