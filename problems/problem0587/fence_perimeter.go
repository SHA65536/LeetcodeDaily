package problem0587

import "sort"

/*
You are given an array trees where trees[i] = [xi, yi] represents the location of a tree in the garden.
You are asked to fence the entire garden using the minimum length of rope as it is expensive.
The garden is well fenced only if all the trees are enclosed.
Return the coordinates of trees that are exactly located on the fence perimeter.
*/

func compareTrees(p1, p2, p3 []int) int {
	var x1, y1 = p1[0], p1[1]
	var x2, y2 = p2[0], p2[1]
	var x3, y3 = p3[0], p3[1]
	return (y3-y2)*(x2-x1) - (y2-y1)*(x3-x2)
}

func outerTrees(trees [][]int) [][]int {
	var upper, lower [][]int
	var treeMap = map[[2]int]struct{}{}
	// Sorting trees by x and y
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0] == trees[j][0] {
			return trees[i][1] < trees[j][1]
		}
		return trees[i][0] < trees[j][0]
	})

	// Building hull
	for i := range trees {
		for len(lower) >= 2 && compareTrees(lower[len(lower)-2], lower[len(lower)-1], trees[i]) > 0 {
			lower = lower[:len(lower)-1]
		}
		for len(upper) >= 2 && compareTrees(upper[len(upper)-2], upper[len(upper)-1], trees[i]) < 0 {
			upper = upper[:len(upper)-1]
		}
		lower = append(lower, trees[i])
		upper = append(upper, trees[i])
	}
	// Removing duplicate points
	for i := range lower {
		treeMap[[2]int{lower[i][0], lower[i][1]}] = struct{}{}
	}
	for i := range upper {
		treeMap[[2]int{upper[i][0], upper[i][1]}] = struct{}{}
	}
	trees = make([][]int, len(treeMap))
	var i = 0
	// Converting unique points into slice
	for k := range treeMap {
		trees[i] = []int{k[0], k[1]}
		i++
	}
	return trees
}
