package problem1061

/*
You are given two strings of the same length s1 and s2 and a string baseStr.
We say s1[i] and s2[i] are equivalent characters.
    For example, if s1 = "abc" and s2 = "cde", then we have 'a' == 'c', 'b' == 'd', and 'c' == 'e'.
Equivalent characters follow the usual rules of any equivalence relation:
    Reflexivity: 'a' == 'a'.
    Symmetry: 'a' == 'b' implies 'b' == 'a'.
    Transitivity: 'a' == 'b' and 'b' == 'c' implies 'a' == 'c'.
For example, given the equivalency information from s1 = "abc" and s2 = "cde", "acd" and "aab" are
equivalent strings of baseStr = "eed", and "aab" is the lexicographically smallest equivalent string of baseStr.
Return the lexicographically smallest equivalent string of baseStr by using the equivalency information from s1 and s2.
*/

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	var res = make([]byte, len(baseStr))
	var graph [26]int        // graph[i] is the parent of i
	var find func(int) int   // Finds the root parent of the input, and fixes the graph while doing it
	var union func(int, int) // Merges two letter's graphs together

	find = func(i int) int {
		// If i has no parent, return itself
		if graph[i] == -1 {
			return i
		}
		// If i has a parent, return the parent's parent
		graph[i] = find(graph[i])
		return graph[i]
	}

	union = func(i, j int) {
		// Getting the root parents
		i, j = find(i), find(j)
		// Updating the graph according to the lexicoragphical order
		if i != j {
			graph[max(i, j)] = min(i, j)
		}
	}

	// Initializing the graph empty
	for i := range graph {
		graph[i] = -1
	}

	// Building graph
	for i := range s1 {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	// Building result
	for i := range baseStr {
		res[i] = byte(find(int(baseStr[i]-'a')) + 'a')
	}

	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
