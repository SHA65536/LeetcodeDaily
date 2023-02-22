package problem1011

/*
A conveyor belt has packages that must be shipped from one port to another within days days.
The ith package on the conveyor belt has a weight of weights[i].
Each day, we load the ship with packages on the conveyor belt (in the order given by weights).
We may not load more weight than the maximum weight capacity of the ship.
Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within days days.
*/

func shipWithinDays(weights []int, days int) int {
	// minw is the minimum cap of the ship
	// maxw is the maximum cap of the ship
	var minw, maxw int
	// The heaviest weight is the minimum capacity
	// and the sum of weights is the maximum capacity (all in one go)
	for _, w := range weights {
		minw = max(minw, w)
		maxw += w
	}
	// Binary search for the minimal weight
	for minw < maxw {
		// need is the days needed for the shipment
		// cur is the current weight of the ship while we load it
		var need, cur int = 1, 0
		// mid is our guess of the weight
		mid := (minw + maxw) / 2
		// Start loading weights
		for _, w := range weights {
			// If we reached capacity, roll over one day
			if cur+w > mid {
				need++
				cur = 0
			}
			// Add weight to the ship
			cur += w
		}
		if need > days {
			// If we exceeded the needed days, we should make the ship
			// capacity larger
			minw = mid + 1
		} else {
			// If we had days to spare, we should make the ship
			// capacity smaller
			maxw = mid
		}
	}
	return minw
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
