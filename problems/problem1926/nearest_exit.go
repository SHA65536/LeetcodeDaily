package problem1926

/*
You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.
In one step, you can move one cell up, down, left, or right. You cannot step into a cell with a wall, and you cannot step outside the maze.
Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze.
The entrance does not count as an exit.
Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.
*/

// Allowed directions for moving
var directions = [4][2]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func nearestExit(maze [][]byte, entrance []int) int {
	var cur, nxt [][2]int // Our BFS layers
	var count int         // Number of moves
	cur = make([][2]int, 0, 4)
	for i := range directions { // Checking around the entrance
		// This is done separately since the entrance does not count as an exit.
		if !isBlock(maze, entrance[0]+directions[i][0], entrance[1]+directions[i][1]) {
			cur = append(cur, [2]int{entrance[0] + directions[i][0], entrance[1] + directions[i][1]})
			maze[entrance[0]+directions[i][0]][entrance[1]+directions[i][1]] = '+' // Marking  as visited
		}
	}
	maze[entrance[0]][entrance[1]] = '+' // Marking  as visited
	// As long as we have options, continue searching
	for len(cur) > 0 {
		nxt = make([][2]int, 0) // Cells to visit in next step
		for i := range cur {
			if isEnd(maze, cur[i][0], cur[i][1]) {
				// If we found the exit, return
				return count + 1
			}
			for j := range directions {
				// Check each direction
				if !isBlock(maze, cur[i][0]+directions[j][0], cur[i][1]+directions[j][1]) {
					// Add to next visit if we can
					nxt = append(nxt, [2]int{cur[i][0] + directions[j][0], cur[i][1] + directions[j][1]})
					maze[cur[i][0]+directions[j][0]][cur[i][1]+directions[j][1]] = '+' // Marking  as visited
				}
			}
			maze[cur[i][0]][cur[i][1]] = '+' // Marking  as visited
		}
		count++   // Didn't find the exit this move
		cur = nxt // Switch layers
	}
	return -1 // If we didn't find an exit, return -1
}

// isEnd checks if cells is outside the maze (an exit)
func isEnd(maze [][]byte, x, y int) bool {
	return x == 0 || y == 0 || x == len(maze)-1 || y == len(maze[0])-1
}

// isBlock checks if cell is not valid or a block
func isBlock(maze [][]byte, x, y int) bool {
	if x < 0 || x >= len(maze) || y < 0 || y >= len(maze[0]) {
		return true
	}
	return maze[x][y] == '+'
}
