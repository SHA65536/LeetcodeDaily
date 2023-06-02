package problem2101

/*
You are given a list of bombs. The range of a bomb is defined as the area where its effect can be felt.
This area is in the shape of a circle with the center as the location of the bomb.
The bombs are represented by a 0-indexed 2D integer array bombs where
bombs[i] = [xi, yi, ri]. xi and yi denote the X-coordinate and Y-coordinate of the location of the ith bomb, whereas ri denotes the radius of its range.
You may choose to detonate a single bomb. When a bomb is detonated, it will detonate all bombs that lie in its range.
These bombs will further detonate the bombs that lie in their ranges.
Given the list of bombs, return the maximum number of bombs that can be detonated if you are allowed to detonate only one bomb.
*/

func maximumDetonation(bombs [][]int) int {
	var res, count int
	var detonated []bool
	// dfs searched for the total number of bombs
	// if we detonate bomb[idx]
	var dfs func(idx int)

	dfs = func(idx int) {
		// Increase the count
		count++
		// Mark as detonated
		detonated[idx] = true
		for i := range bombs {
			// If in range, detonate
			if !detonated[i] && inRange(bombs[idx], bombs[i]) {
				detonated[i] = true
				dfs(i)
			}
		}
	}

	// Try all bombs
	for i := range bombs {
		// Clean slate
		count = 0
		detonated = make([]bool, len(bombs))
		dfs(i)
		res = max(res, count)
	}

	return res
}

// inRange shows if dst bomb is in range of dstBomb
func inRange(srcBomb, dstBomb []int) bool {
	var dx, dy = srcBomb[0] - dstBomb[0], srcBomb[1] - dstBomb[1]
	return dx*dx+dy*dy <= srcBomb[2]*srcBomb[2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
