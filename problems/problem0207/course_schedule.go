package problem0207

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that
you must take course bi first if you want to take course ai.
For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	var res int
	// cur and nxt are for the BFS
	var cur, nxt []int

	// preCnt[x] shows how many prerequisites x has
	var preCnt = make([]int, numCourses)
	// depends[x] shows the courses that depend on x
	var depends = make([][]int, numCourses)

	// Parsing prerequisites list
	for _, p := range prerequisites {
		depends[p[1]] = append(depends[p[1]], p[0])
		preCnt[p[0]]++
	}

	// Finding all the courses that don't have prerequisites
	for i := range preCnt {
		if preCnt[i] == 0 {
			cur = append(cur, i)
		}
	}

	// BFS over the courses that have no prerequisites
	for len(cur) > 0 {
		for _, c := range cur {
			// Mark the course as done
			res++
			// Look at all the courses that depend on it
			for _, n := range depends[c] {
				// Remove the dependencies since we've completed the course
				preCnt[n]--
				// If it has no additional dependencies
				if preCnt[n] == 0 {
					// Add it to be completed
					nxt = append(nxt, n)
				}
			}
		}
		// Switch BFS layers
		cur, nxt = nxt, cur
		nxt = nxt[:0]
	}

	return res == numCourses
}
