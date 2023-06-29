package problem0864

/*
You are given an m x n grid grid where:
    '.' is an empty cell.
    '#' is a wall.
    '@' is the starting point.
    Lowercase letters represent keys.
    Uppercase letters represent locks.
You start at the starting point and one move consists of walking one space in one of the four cardinal directions.
You cannot walk outside the grid, or walk into a wall.
If you walk over a key, you can pick it up and you cannot walk over a lock unless you have its corresponding key.
For some 1 <= k <= 6, there is exactly one lowercase and one uppercase letter of the first k letters of the English alphabet in the grid.
This means that there is exactly one key for each lock, and one lock for each key;
and also that the letters used to represent the keys and locks were chosen in the same order as the English alphabet.
Return the lowest number of moves to acquire all keys. If it is impossible, return -1.
*/

var moves = [4][2]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

type Path struct {
	X, Y int
	Keys uint32
}

func shortestPathAllKeys(grid []string) int {
	var res int
	sx, sy, dst := FindStartAndKeys(grid)
	var cur, nxt []Path
	cur = []Path{{sx, sy, 0}}
	var cache = makeCache(grid)
	for len(cur) > 0 {
		nxt = nxt[:0]
		res++
		for _, c := range cur {
			if c.Keys == dst {
				return res - 2
			}
			nxt = append(nxt, NextCells(grid, cache, c)...)
		}
		cur, nxt = nxt, cur
	}
	return -1
}

func makeCache(grid []string) [][]map[uint32]bool {
	var res = make([][]map[uint32]bool, len(grid))
	for i := range grid {
		res[i] = make([]map[uint32]bool, len(grid[0]))
		for j := range grid[i] {
			res[i][j] = map[uint32]bool{}
		}
	}
	return res
}

func NextCells(grid []string, cache [][]map[uint32]bool, cur Path) []Path {
	var res = make([]Path, 0, 4)
	if isOutside(grid, cur.X, cur.Y) {
		return nil
	}
	if cache[cur.X][cur.Y][cur.Keys] {
		return nil
	}
	if grid[cur.X][cur.Y] == '#' {
		return nil
	}
	if isUpper(grid[cur.X][cur.Y]) && !cur.HasKey(grid[cur.X][cur.Y]) {
		return nil
	}
	if isLower(grid[cur.X][cur.Y]) {
		cur.AddKey(grid[cur.X][cur.Y])
	}
	cache[cur.X][cur.Y][cur.Keys] = true
	for _, move := range moves {
		if !isOutside(grid, cur.X, cur.Y) {
			res = append(res, Path{cur.X + move[0], cur.Y + move[1], cur.Keys})
		}
	}
	return res
}

func isOutside(grid []string, x, y int) bool {
	return x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0])
}

// FindStartAndKeys finds the starting position and the desired keys
func FindStartAndKeys(grid []string) (int, int, uint32) {
	var sx, sy int
	var key Path
	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case '@':
				sx, sy = i, j
			case '.', '#':
			default:
				if isLower(grid[i][j]) {
					key.AddKey(grid[i][j])
				}
			}
		}
	}
	return sx, sy, key.Keys
}

func isUpper(letter byte) bool {
	return letter >= 'A' && letter <= 'Z'
}

func isLower(letter byte) bool {
	return letter >= 'a' && letter <= 'z'
}

func (p *Path) AddKey(key byte) {
	p.Keys |= 1 << (key - 'a')
}

func (p *Path) HasKey(lock byte) bool {
	var bit uint32 = 1 << (lock - 'A')
	return (p.Keys & bit) == bit
}
