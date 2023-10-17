package problem1361

/*
You have n binary tree nodes numbered from 0 to n - 1 where node i has two children leftChild[i] and rightChild[i],
return true if and only if all the given nodes form exactly one valid binary tree.
If node i has no left child then leftChild[i] will equal -1, similarly for the right child.
Note that the nodes have no values and that we only use the node numbers in this problem.
*/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// uf[i] represents the group that is the parent of group i
	var uf = make([]int, n)
	// indegree[i] represents how many nodes lead to node i
	var indegree = make([]int, n)

	// Initialize the disjoint set (union find) group
	for i := range uf {
		uf[i] = i
	}

	// Loop over the nodes and build the graph
	for i := 0; i < n; i++ {
		// If the is a child node
		if leftChild[i] != -1 {
			// Increment the node's indegree
			indegree[leftChild[i]]++
			// Update the parent graph
			union(uf, leftChild[i], i)
		}

		// Same for right child
		if rightChild[i] != -1 {
			indegree[rightChild[i]]++
			union(uf, rightChild[i], i)
		}
	}

	// There should be only 1 root (a node with no indegree)
	var foundRoot = false
	for i := range indegree {
		if indegree[i] > 1 {
			return false
		}
		if indegree[i] == 0 {
			if foundRoot {
				return false
			}
			foundRoot = true
		}
	}

	// Find how many different groups there are
	var uniqGroups = map[int]bool{}
	for i := range uf {
		uniqGroups[find(uf, i)] = true
	}

	// There should be only one group and one root
	return len(uniqGroups) == 1 && foundRoot
}

func find(uf []int, x int) int {
	if uf[x] == x {
		// If x is the parent of itself, it is the root of the group
		return uf[x]
	} else {
		// If x is not the parent of itself, we call this function again
		// to find the real parent, and update the map
		uf[x] = find(uf, uf[x])
		return uf[x]
	}
}

func union(uf []int, x, y int) {
	var rootx, rooty int
	// Finding the roots of x and y
	rootx = find(uf, x)
	rooty = find(uf, y)
	// Setting the root of rootx be rooty effectivly merging the groups
	uf[rootx] = rooty
}
