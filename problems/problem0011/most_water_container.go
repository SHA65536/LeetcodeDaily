package problem0011

/*
https://leetcode.com/problems/container-with-most-water

You are given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

func maxArea(height []int) int {
	var left, right, area, maxarea int
	right = len(height) - 1
	for left < right {
		if height[left] > height[right] {
			area = height[right] * (right - left)
			right--
		} else {
			area = height[left] * (right - left)
			left++
		}
		if area > maxarea {
			maxarea = area
		}
	}
	return maxarea
}
