package problem1496

/*
Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively.
You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.
Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited.
Return false otherwise.
*/

type Pos struct {
	X, Y int
}

func isPathCrossing(path string) bool {
	var cur Pos
	var cache = map[Pos]struct{}{cur: {}}
	for i := range path {
		switch path[i] {
		case 'N':
			cur.X++
		case 'S':
			cur.X--
		case 'E':
			cur.Y++
		case 'W':
			cur.Y--
		}
		if _, ok := cache[cur]; ok {
			return true
		}
		cache[cur] = struct{}{}
	}
	return false
}
