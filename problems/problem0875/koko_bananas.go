package problem0875

/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.
Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
Return the minimum integer k such that she can eat all the bananas within h hours.
*/

// We employ binary search here to determine the optimal speed
func minEatingSpeed(piles []int, h int) int {
	var left, right int = 1, max(piles) * len(piles)
	for left < right {
		// We pick a speed in the middle
		mid := (left + right) / 2
		if canEat(piles, h, mid) {
			// If koko can eat all the bananas with this speed
			// we check a lower speed
			right = mid
		} else {
			// If koko can't eat all the bananasw ith this speed
			// we check a higher speed
			left = mid + 1
		}
	}
	return left
}

// canEat returns true if koko can eat all the bananas
// in h hours with a speed of s bananas per hour
func canEat(piles []int, h int, s int) bool {
	var time int
	for i := 0; i < len(piles) && time <= h; i++ {
		time += piles[i] / s
		if piles[i]%s != 0 {
			time++
		}
	}
	return time <= h
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
