package problem0733

/*
An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill,
consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
Replace the color of all of the aforementioned pixels with newColor.

Return the modified image after performing the flood fill.
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited := make([][]bool, len(image))
	for i := 0; i < len(image); i++ {
		visited[i] = make([]bool, len(image[0]))
	}
	floodFillHelper(image, visited, sr, sc, image[sr][sc], newColor)
	return image
}

func floodFillHelper(image [][]int, visited [][]bool, curR, curC, oldC, newC int) {
	if curC < 0 || curR < 0 || curR >= len(image) || curC >= len(image[0]) {
		return
	}
	if visited[curR][curC] {
		return
	}
	visited[curR][curC] = true
	if image[curR][curC] == oldC {
		image[curR][curC] = newC
		floodFillHelper(image, visited, curR+1, curC, oldC, newC)
		floodFillHelper(image, visited, curR-1, curC, oldC, newC)
		floodFillHelper(image, visited, curR, curC+1, oldC, newC)
		floodFillHelper(image, visited, curR, curC-1, oldC, newC)
	}
}
