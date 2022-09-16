package problem0987

import (
	. "leetcodedaily/helpers/binarytree"
	"sort"
)

/*
Given the root of a binary tree, calculate the vertical order traversal of the binary tree.
For each node at position (row, col), its left and right children will be at positions (row + 1, col - 1) and (row + 1, col + 1) respectively.
The root of the tree is at (0, 0).
The vertical order traversal of a binary tree is a list of top-to-bottom orderings
for each column index starting from the leftmost column and ending on the rightmost column.
There may be multiple nodes in the same row and same column. In such a case, sort these nodes by their values.
Return the vertical order traversal of the binary tree.
*/

func verticalTraversal(root *TreeNode) [][]int {
	var results = [][]int{}
	var resmap = map[int][][2]int{}
	var min int
	solveTree(root, &min, 0, 0, resmap)
	for i := min; true; i++ {
		if val, ok := resmap[i]; ok {
			var temp = []int{}
			for _, pair := range val {
				temp = append(temp, pair[1])
			}
			results = append(results, temp)
		} else {
			break
		}
	}
	return results
}

func solveTree(root *TreeNode, min *int, curX, curY int, resmap map[int][][2]int) {
	if root == nil {
		return
	}

	if *min > curX {
		*min = curX
	}

	solveTree(root.Left, min, curX-1, curY-1, resmap)

	if _, ok := resmap[curX]; !ok {
		resmap[curX] = [][2]int{{curY, root.Val}}
	} else {
		resmap[curX] = append(resmap[curX], [2]int{curY, root.Val})
		sort.Slice(resmap[curX], func(i, j int) bool {
			if resmap[curX][i][0] == resmap[curX][j][0] {
				return resmap[curX][i][1] < resmap[curX][j][1]
			}
			return resmap[curX][i][0] > resmap[curX][j][0]
		})
	}

	solveTree(root.Right, min, curX+1, curY-1, resmap)
}
