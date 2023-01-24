package problem0909

/*
You are given an n x n integer matrix board where the cells are labeled from 1 to n2 in a Boustrophedon style starting from the bottom left of the board
(i.e. board[n - 1][0]) and alternating direction each row.
You start on square 1 of the board. In each move, starting from square curr, do the following:
    Choose a destination square next with a label in the range [curr + 1, min(curr + 6, n2)].
    This choice simulates the result of a standard 6-sided die roll: i.e., there are always at most 6 destinations, regardless of the size of the board.
    If next has a snake or ladder, you must move to the destination of that snake or ladder. Otherwise, you move to next.
    The game ends when you reach the square n2.
A board square on row r and column c has a snake or ladder if board[r][c] != -1.
The destination of that snake or ladder is board[r][c]. Squares 1 and n2 do not have a snake or ladder.
Note that you only take a snake or ladder at most once per move.
If the destination to a snake or ladder is the start of another snake or ladder, you do not follow the subsequent snake or ladder.
For example, suppose the board is [[-1,4],[-1,3]], and on the first move, your destination square is 2.
You follow the ladder to square 3, but do not follow the subsequent ladder to 4.
Return the least number of moves required to reach the square n2. If it is not possible to reach the square, return -1.
*/

func snakesAndLadders(board [][]int) int {
	// arr is the board squashed into 1 dimension
	var arr = convertBoard(board)
	// visited[i] represents true if we visited this spot, otherwise false
	var visited = make([]bool, len(arr))
	// cur is the current batch of the BFS, nxt is the future
	var cur, nxt []int
	var res int
	cur = []int{0}

	for len(cur) > 0 {
		res++
		nxt = []int{}
		for i := range cur {
			visited[cur[i]] = true
			// If we reached the end, return current steps
			if cur[i] == len(arr)-1 {
				return res - 1
			}
			// Add all the possible next moves to nxt
			for j := 1; j <= 6 && cur[i]+j < len(arr); j++ {
				var target int = cur[i] + j
				// If target is a ladder, move with the ladder
				if arr[target] != -1 {
					target = arr[target] - 1
				}
				// Add to nxt if not visited
				if !visited[target] {
					nxt = append(nxt, target)
				}
			}
		}
		// Next batch
		cur = nxt
	}

	// If we reached here, we didn't find the target
	return -1
}

func convertBoard(board [][]int) []int {
	var res = make([]int, len(board)*len(board))
	var cur int

	for x := len(board) - 1; x >= 0; x-- {
		if (x%2 == 0) != (len(board)%2 == 0) {
			for y := 0; y < len(board); y++ {
				res[cur] = board[x][y]
				cur++
			}
		} else {
			for y := len(board) - 1; y >= 0; y-- {
				res[cur] = board[x][y]
				cur++
			}
		}
	}
	return res
}

// Here is my attempt and DFS, better to do a bfs.
/*
const (
	EmptySquare = -1
	NotVisited  = -1
	Visited     = -2
)

func snakesAndLadders(board [][]int) int {
	var res int
	// arr is the board squashed into 1 dimension
	var arr = convertBoard(board)
	// dp is cache of what we calculated,
	// dp[i] is the number of moves to win from square i
	var dp = make([]int, len(arr))
	// dfs(i) calculates the number of moves to win from square i
	// updates cache accordingly
	var dfs func(int) int

	for i := range dp {
		dp[i] = NotVisited
	}
	// The end takes 0 steps to win
	dp[len(dp)-1] = 0

	dfs = func(cur int) int {
		var minMoves int = 1000
		// If we already calculated this position, return cache
		if dp[cur] >= 0 {
			return dp[cur]
		}

		// If we visited already, don't loop
		if dp[cur] == Visited {
			return minMoves
		}

		// Mark as visited
		dp[cur] = Visited

		for i := 1; i <= 6 && i+cur < len(arr); i++ {
			var temp int
			if arr[cur+i] != EmptySquare {
				temp = dfs(arr[cur+i] - 1)
			} else {
				temp = dfs(cur + i)
			}
			if temp < minMoves {
				minMoves = temp
			}
		}

		// Unmark as visited
		dp[cur] = minMoves + 1
		return minMoves + 1
	}

	res = dfs(0)
	// In case it's impossible
	if res >= 1000 {
		return -1
	}
	return res
}
*/
