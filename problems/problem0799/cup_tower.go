package problem0799

/*
We stack glasses in a pyramid, where the first row has 1 glass, the second row has 2 glasses, and so on until the 100th row.
Each glass holds one cup of champagne.
Then, some champagne is poured into the first glass at the top.
When the topmost glass is full, any excess liquid poured will fall equally to the glass immediately to the left and right of it.
When those glasses become full, any excess champagne will fall equally to the left and right of those glasses, and so on.
(A glass at the bottom row has its excess champagne fall on the floor.)
For example, after one cup of champagne is poured, the top most glass is full.
After two cups of champagne are poured, the two glasses on the second row are half full.
After three cups of champagne are poured, those two cups become full - there are 3 full glasses total now.
After four cups of champagne are poured, the third row has the middle glass half full,
and the two outside glasses are a quarter full, as pictured below.
Now after pouring some non-negative integer cups of champagne,
return how full the jth glass in the ith row is (both i and j are 0-indexed.)
*/

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var cur, prev = make([]float64, 2, query_row+3), make([]float64, 1, query_row+3)
	prev[0] = float64(poured)
	for r := 1; r <= query_row+1; r++ {
		cur[0] = max(0, (prev[0]-1)/2)
		for i := 1; i < len(cur)-1; i++ {
			cur[i] = max(0, (prev[i-1]-1)/2)
			cur[i] += max(0, (prev[i]-1)/2)
		}
		cur[len(cur)-1] = max(0, (prev[len(prev)-1]-1)/2)
		prev, cur = cur, prev[:r+2]
	}
	return min(1, cur[query_glass])
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
