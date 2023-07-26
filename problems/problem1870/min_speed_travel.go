package problem1870

import "math"

/*
You are given a floating-point number hour, representing the amount of time you have to reach the office.
To commute to the office, you must take n trains in sequential order.
You are also given an integer array dist of length n, where dist[i] describes the distance (in kilometers) of the ith train ride.
Each train can only depart at an integer hour, so you may need to wait in between each train ride.
    For example, if the 1st train ride takes 1.5 hours,
	you must wait for an additional 0.5 hours before you can depart on the 2nd train ride at the 2 hour mark.
Return the minimum positive integer speed (in kilometers per hour) that all the
trains must travel at for you to reach the office on time, or -1 if it is impossible to be on time.
Tests are generated such that the answer will not exceed 107 and hour will have at most two digits after the decimal point.
*/

func minSpeedOnTime(dist []int, hour float64) int {
	// Each train has to take atleast 1 hour
	// so if we have more trains than hours we cannot get in time
	if float64(len(dist)-1) >= hour {
		return -1
	}
	var low, high = 0, 10000000
	// Simple binary search
	for low < high {
		mid := low + (high-low)/2
		// We calculate if we can get in time with mid speed
		if calcTime(dist, mid, hour) {
			// If we can, we should consider a lower speed
			high = mid
		} else {
			// If we can't, we need higher speed
			low = mid + 1
		}
	}
	return low
}

func calcTime(dist []int, speed int, hour float64) bool {
	var curTime float64
	for i := 0; i < len(dist) && curTime < hour; i++ {
		curTime = math.Ceil(curTime)
		curTime += float64(dist[i]) / float64(speed)
	}
	return curTime <= hour
}
