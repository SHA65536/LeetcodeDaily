package problem0999

/*
On an 8 x 8 chessboard, there is exactly one white rook 'R' and some number of white bishops 'B', black pawns 'p', and empty squares '.'.
When the rook moves, it chooses one of four cardinal directions (north, east, south, or west),
then moves in that direction until it chooses to stop, reaches the edge of the board, captures a black pawn, or is blocked by a white bishop.
A rook is considered attacking a pawn if the rook can capture the pawn on the rook's turn.
The number of available captures for the white rook is the number of pawns that the rook is attacking.
Return the number of available captures for the white rook.
*/

func numRookCaptures(board [][]byte) int {
	var rX, rY, res int
	// Finding the rook
findRook:
	for rX = 0; rX < len(board); rX++ {
		for rY = 0; rY < len(board[0]); rY++ {
			if board[rX][rY] == 'R' {
				break findRook
			}
		}
	}
	// Checking South
	for nX := rX + 1; nX < len(board); nX++ {
		if board[nX][rY] == 'B' {
			break
		} else if board[nX][rY] == 'p' {
			res++
			break
		}
	}
	// Checking North
	for nX := rX - 1; nX >= 0; nX-- {
		if board[nX][rY] == 'B' {
			break
		} else if board[nX][rY] == 'p' {
			res++
			break
		}
	}
	// Checking East
	for nY := rY + 1; nY < len(board); nY++ {
		if board[rX][nY] == 'B' {
			break
		} else if board[rX][nY] == 'p' {
			res++
			break
		}
	}
	// Checkin West
	for nY := rY - 1; nY >= 0; nY-- {
		if board[rX][nY] == 'B' {
			break
		} else if board[rX][nY] == 'p' {
			res++
			break
		}
	}
	return res
}
