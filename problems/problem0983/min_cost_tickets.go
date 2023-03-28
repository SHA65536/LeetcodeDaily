package problem0983

/*
You have planned some train traveling one year in advance.
The days of the year in which you will travel are given as an integer array days.
Each day is an integer from 1 to 365.
Train tickets are sold in three different ways:
    a 1-day pass is sold for costs[0] dollars,
    a 7-day pass is sold for costs[1] dollars, and
    a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.
    For example, if we get a 7-day pass on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.
Return the minimum number of dollars you need to travel every day in the given list of days.
*/

func mincostTickets(days []int, costs []int) int {
	// cache stores the costs for the last 30 days of travel
	var cache [30]int
	// travel[i] tells us if we travel on the ith day
	var travel = make(map[int]struct{}, len(days))
	var l = len(days) - 1
	for _, d := range days {
		travel[d] = struct{}{}
	}

	// Loop over all calendar days (not the given days)
	for i := 1; i <= 365; i++ {
		if _, ok := travel[i]; !ok {
			// If we don't travel this day, take the last day cost
			cache[i%30] = cache[(i-1)%30]
		} else {
			cache[i%30] = min(
				cache[(i-1)%30]+costs[0],        // 1 day ticket
				cache[max(0, i-7)%30]+costs[1],  // 7 day ticket
				cache[max(0, i-30)%30]+costs[2], // 30 day ticket
			)

		}
	}

	return cache[days[l]%30]
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
