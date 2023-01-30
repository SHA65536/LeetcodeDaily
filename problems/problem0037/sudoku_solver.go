package problem0037

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.
A sudoku solution must satisfy all of the following rules:
    Each of the digits 1-9 must occur exactly once in each row.
    Each of the digits 1-9 must occur exactly once in each column.
    Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.
*/

const (
	One  byte = '1'
	Nine byte = '9'
	Open byte = '.'
)

type Soduku struct {
	Board [][]byte
}

func (s *Soduku) Solve() bool {
	// Finding open spots
	oi, oj, ok := s.FindOpen()
	if !ok { // If not spots, we solved
		return true
	}
	// Trying to put numbers
	for i := One; i <= Nine; i++ {
		s.Board[oi][oj] = i
		if s.IsGood(oi, oj) && s.Solve() {
			return true
		}
	}
	s.Board[oi][oj] = Open
	return false
}

func (s *Soduku) IsGood(i, j int) bool {
	var row [9]bool
	var col [9]bool
	// Checking row and col
	for ci := range row {
		if s.Board[i][ci] != Open {
			if row[s.Board[i][ci]-One] {
				return false
			}
			row[s.Board[i][ci]-One] = true
		}
		if s.Board[ci][j] != Open {
			if col[s.Board[ci][j]-One] {
				return false
			}
			col[s.Board[ci][j]-One] = true
		}
	}
	// Checking square
	row = [9]bool{}
	r, c := (i/3)*3, (j/3)*3
	for ci := 0; ci < 3; ci++ {
		for cj := 0; cj < 3; cj++ {
			if s.Board[r+ci][c+cj] == Open {
				continue
			}
			if row[s.Board[r+ci][c+cj]-One] {
				return false
			}
			row[s.Board[r+ci][c+cj]-One] = true
		}
	}
	return true
}

func (s *Soduku) FindOpen() (int, int, bool) {
	for i := range s.Board {
		for j := range s.Board {
			if s.Board[i][j] == Open {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

func solveSudoku(board [][]byte) {
	var s = &Soduku{Board: board}
	s.Solve()
}
