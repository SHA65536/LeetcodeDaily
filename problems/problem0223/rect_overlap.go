package problem0223

/*
Given the coordinates of two rectilinear rectangles in a 2D plane, return the total area covered by the two rectangles.
The first rectangle is defined by its bottom-left corner (ax1, ay1) and its top-right corner (ax2, ay2).
The second rectangle is defined by its bottom-left corner (bx1, by1) and its top-right corner (bx2, by2).
*/

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	xOverlap := max(min(ax2, bx2)-max(ax1, bx1), 0)
	yOverlap := max(min(ay2, by2)-max(ay1, by1), 0)
	return area(ax1, ay1, ax2, ay2) + area(bx1, by1, bx2, by2) - xOverlap*yOverlap
}

func area(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
