package problem0138

/*
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station.
You begin the journey with an empty tank at one of the gas stations.
Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction
otherwise return -1. If there exists a solution, it is guaranteed to be unique
*/

func canCompleteCircuit(gas []int, cost []int) int {
	var total, tank, start int
	for i := range gas {
		// Total gas to see if its possible to complete
		total += gas[i] - cost[i]
		// Current tank
		tank += gas[i] - cost[i]
		// If we failed in going around
		if tank < 0 {
			// Reset the tank
			tank = 0
			start = i + 1
		}
	}
	// If there's enough gas to make a loop
	if total >= 0 {
		return start
	}
	return -1
}
