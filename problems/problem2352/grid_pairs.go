package problem2352

import (
	"reflect"
	"unsafe"
)

/*
Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.
A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).
*/

func equalPairsLoops(grid [][]int) int {
	var res int
	for r := range grid {
		for c := range grid {
			var m int
			for ; m < len(grid) && grid[r][m] == grid[m][c]; m++ {
			}
			if m == len(grid) {
				res++
			}
		}
	}
	return res
}

func equalPairsMap(grid [][]int) int {
	var res int
	var hashMap = make(map[string]int, len(grid))
	for r := range grid {
		hashMap[sliceToString(grid[r])]++
	}
	for c := range grid {
		var temp = make([]int, len(grid))
		for i := range grid[c] {
			temp[i] = grid[i][c]
		}
		res += hashMap[sliceToString(temp)]
	}
	return res
}

func sliceToString(slice []int) string {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&slice))
	byteSlice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: header.Data,
		Len:  header.Len * int(unsafe.Sizeof(slice[0])),
		Cap:  header.Cap * int(unsafe.Sizeof(slice[0])),
	}))
	return *(*string)(unsafe.Pointer(&byteSlice))
}
