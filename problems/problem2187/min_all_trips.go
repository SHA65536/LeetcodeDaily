package problem2187

/*
You are given an array time where time[i] denotes the time taken by the ith bus to complete one trip.
Each bus can make multiple trips successively; that is, the next trip can start immediately after completing the current trip.
Also, each bus operates independently; that is, the trips of one bus do not influence the trips of any other bus.
You are also given an integer totalTrips, which denotes the number of trips all buses should make in total.
Return the minimum time required for all buses to complete at least totalTrips trips.
*/

func minimumTime(time []int, totalTrips int) int64 {
	var left, right int64 = 1, int64(max(time)) * int64(totalTrips) * 10
	for left < right {
		mid := left + (right-left)/2
		if !calcTime(mid, totalTrips, time) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func calcTime(a int64, total int, times []int) bool {
	var sum int64
	totalLong := int64(total)
	for _, t := range times {
		sum += a / int64(t)
		if sum >= totalLong {
			break
		}
	}
	return sum >= totalLong
}

func max(in []int) int {
	var res = in[0]
	for i := range in {
		if in[i] > res {
			res = in[i]
		}
	}
	return res
}
