package problem0427

import . "leetcodedaily/helpers/quadtree"

/*
Given a n * n matrix grid of 0's and 1's only. We want to represent the grid with a Quad-Tree.
Return the root of the Quad-Tree representing the grid.
Notice that you can assign the value of a node to True or False when isLeaf is False, and both are accepted in the answer.
A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:
    val: True if the node represents a grid of 1's or False if the node represents a grid of 0's.
    isLeaf: True if the node is leaf node on the tree or False if the node has the four children.
class Node {
    public boolean val;
    public boolean isLeaf;
    public Node topLeft;
    public Node topRight;
    public Node bottomLeft;
    public Node bottomRight;
}
We can construct a Quad-Tree from a two-dimensional area using the following steps:
    If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
    If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
    Recurse for each of the children with the proper sub-grid.
*/

func construct(grid [][]int) *Node {
	return constructHelper(grid, 0, 0, len(grid))
}

func constructHelper(grid [][]int, x, y, l int) *Node {
	if l == 1 {
		return &Node{Val: grid[x][y] == 1, IsLeaf: true}
	}
	tl := constructHelper(grid, x, y, l/2)
	tr := constructHelper(grid, x, y+l/2, l/2)
	bl := constructHelper(grid, x+l/2, y, l/2)
	br := constructHelper(grid, x+l/2, y+l/2, l/2)
	if tl.IsLeaf && All(tl.IsLeaf, tr.IsLeaf, bl.IsLeaf, br.IsLeaf) && All(tl.Val, tr.Val, bl.Val, br.Val) {
		return &Node{Val: tl.Val, IsLeaf: true}
	} else {
		return &Node{
			Val: true, IsLeaf: false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}
}

func All[T comparable](args ...T) bool {
	var test T = args[0]
	for i := range args {
		if args[i] != test {
			return false
		}
	}
	return true
}

func constructNoGeneric(grid [][]int) *Node {
	return constructHelperNoGeneric(grid, 0, 0, len(grid))
}

func constructHelperNoGeneric(grid [][]int, x, y, l int) *Node {
	if l == 1 {
		return &Node{Val: grid[x][y] == 1, IsLeaf: true}
	}
	tl := constructHelperNoGeneric(grid, x, y, l/2)
	tr := constructHelperNoGeneric(grid, x, y+l/2, l/2)
	bl := constructHelperNoGeneric(grid, x+l/2, y, l/2)
	br := constructHelperNoGeneric(grid, x+l/2, y+l/2, l/2)
	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
		return &Node{Val: tl.Val, IsLeaf: true}
	} else {
		return &Node{
			Val: true, IsLeaf: false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}
}
