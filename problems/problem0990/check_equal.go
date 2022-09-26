package problem0990

/*
You are given an array of strings equations that represent relationships between variables
where each string equations[i] is of length 4 and takes one of two different forms: "xi==yi" or "xi!=yi".Here, xi and yi are lowercase letters
(not necessarily different) that represent one-letter variable names.
Return true if it is possible to assign integers to variable names so as to satisfy all the given equations, or false otherwise.
*/

const OFFSET = 'a'

func findParent(root []int, x int) int {
	if x != root[x] {
		root[x] = findParent(root, root[x])
	}
	return root[x]
}

func equationsPossible(equations []string) bool {
	var root = make([]int, 26)
	var parentLeft, parentRight int
	// Making each variable it's own parent
	for i := 0; i < 26; i++ {
		root[i] = i
	}
	// Building equalities
	for _, eq := range equations {
		if eq[1] == '=' {
			parentLeft = findParent(root, int(eq[0]-OFFSET))
			parentRight = findParent(root, int(eq[3]-OFFSET))
			root[parentLeft] = parentRight
		}
	}
	// Checking inequalities
	for _, eq := range equations {
		if eq[1] == '!' {
			parentLeft = findParent(root, int(eq[0]-OFFSET))
			parentRight = findParent(root, int(eq[3]-OFFSET))
			if parentLeft == parentRight {
				return false
			}
		}
	}
	return true
}
