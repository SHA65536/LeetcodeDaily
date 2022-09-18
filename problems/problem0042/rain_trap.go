package problem0042

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
*/

func trap(height []int) int {
	var res, left, right, mLeft, mRight int
	right = len(height) - 1
	for left <= right {
		if height[left] <= height[right] {
			// If the current left is the lower boundry of the pool
			if height[left] >= mLeft {
				// If left is a new max height, there's no water above it
				mLeft = height[left]
			} else {
				// If it's not the new max height, it's covered in water
				res += mLeft - height[left]
			}
			left++
		} else {
			// If the current right is the lower boundry of the pool
			if height[right] >= mRight {
				// If right is a new max height, there's no water above it
				mRight = height[right]
			} else {
				// If it's not the new max height, it's covered in water
				res += mRight - height[right]
			}
			right--
		}
	}
	return res
}
