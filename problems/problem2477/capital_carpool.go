package problem2477

/*
There is a tree (i.e., a connected, undirected graph with no cycles) structure country network consisting of n cities numbered from
0 to n - 1 and exactly n - 1 roads. The capital city is city 0.
You are given a 2D integer array roads where roads[i] = [ai, bi] denotes that there exists a bidirectional road connecting cities ai and bi.
There is a meeting for the representatives of each city. The meeting is in the capital city.
There is a car in each city. You are given an integer seats that indicates the number of seats in each car.
A representative can use the car in their city to travel or change the car and ride with another representative.
The cost of traveling between two cities is one liter of fuel.
Return the minimum number of liters of fuel to reach the capital city.
*/

func minimumFuelCost(roads [][]int, seats int) int64 {
	var res int64
	var graph = make([][]int, len(roads)+1)
	var dfs func(cur, prev int) int
	// Building the graph from the edges
	for i := range roads {
		graph[roads[i][0]] = append(graph[roads[i][0]], roads[i][1])
		graph[roads[i][1]] = append(graph[roads[i][1]], roads[i][0])
	}

	// dfs calculates the amount of people that need
	// to get to node cur, and updates the result accordingly
	dfs = func(cur, prev int) int {
		var people = 1
		// Calculate number of people in each adjacent node
		// except the previous node (the one close to 0, since we started at 0)
		for _, nxt := range graph[cur] {
			if nxt == prev {
				continue
			}
			people += dfs(nxt, cur)
		}
		// If we're not at the destination, all of these people need to move
		// So we fit them all into cars, and the amount of cars is added to the answer
		if cur != 0 {
			res += int64((people + seats - 1) / seats)
		}
		// Return amount of people for the rest of the calculation
		return people
	}

	// Start from node 0
	dfs(0, 0)

	return res
}
