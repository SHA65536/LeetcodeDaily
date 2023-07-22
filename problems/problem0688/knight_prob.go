package problem0688

/*
On an n x n chessboard, a knight starts at the cell (row, column) and attempts to make exactly k moves.
The rows and columns are 0-indexed, so the top-left cell is (0, 0), and the bottom-right cell is (n - 1, n - 1).
A chess knight has eight possible moves it can make, as illustrated below.
Each move is two cells in a cardinal direction, then one cell in an orthogonal direction.
Each time the knight is to move, it chooses one of eight possible moves uniformly at random
(even if the piece would go off the chessboard) and moves there.
The knight continues moving until it has made exactly k moves or has moved off the chessboard.
Return the probability that the knight remains on the board after it has stopped moving.
*/

type Coord struct {
	X, Y int
}

var KnightMoves = [8]Coord{
	{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2},
	{1, -2}, {1, 2}, {2, -1}, {2, 1},
}

func knightProbability(n, k, row, col int) float64 {
	// cur and prev show the probability of staying on board
	// for the current and previous moves for each position
	var cur, prev [][]float64 = createBoard(n), createBoard(n)

	for k >= 0 {
		// For each square in the current board
		for i := range cur {
			for j := range cur[i] {
				// Calculate the probability of going out of bounds
				// using the probability from the previous move
				var probSum float64
				for _, delta := range KnightMoves {
					// Sum the probability of all the moves that stay on board
					nxt := Coord{X: i + delta.X, Y: j + delta.Y}
					if legalMove(n, nxt) {
						probSum += prev[nxt.X][nxt.Y]
					}
				}
				// Divide by the number of possible moves (8 for knight)
				cur[i][j] = probSum / 8
			}
		}
		// Switch cur and prev for next move
		cur, prev = prev, cur
		k--
	}

	return cur[row][col]
}

func createBoard(n int) [][]float64 {
	var res = make([][]float64, n)
	for i := range res {
		res[i] = make([]float64, n)
		for j := range res[i] {
			res[i][j] = 1
		}
	}
	return res
}

func legalMove(n int, cur Coord) bool {
	return cur.X >= 0 && cur.Y >= 0 && cur.X < n && cur.Y < n
}

/* Naive brute force, runs out of memory
func knightProbability(n, k, row, col int) float64 {
	// cur and nxt save the current legal moves to be explored
	var cur, nxt []Coord
	// allMoves is the amount of all possible moves
	var allMoves = math.Pow(8, float64(k))
	cur = []Coord{{row, col}}
	// BFS searching for all legal moves after k moves
	for len(cur) > 0 && k > 0 {
		// For each move in the search
		for _, cur_move := range cur {
			// Find all possible moves
			for _, delta_cord := range KnightMoves {
				new_move := Coord{
					X: cur_move.X + delta_cord.X,
					Y: cur_move.Y + delta_cord.Y,
				}
				// If the move is legal, add to next layer of search
				if legalMove(n, new_move) {
					nxt = append(nxt, new_move)
				}
			}
		}
		// Switch BFS layers
		cur, nxt = nxt, cur
		nxt = nxt[:0]
		k--
	}
	// All the legal moves are contained in cur, so we divide by possible moves
	// to show probability of staying on board
	return float64(len(cur)) / allMoves
}*/
