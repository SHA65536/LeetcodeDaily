package problem2050

/*
You are given an integer n, which indicates that there are n courses labeled from 1 to n.
You are also given a 2D integer array relations where relations[j] = [prevCoursej, nextCoursej] denotes that
course prevCoursej has to be completed before course nextCoursej (prerequisite relationship).
Furthermore, you are given a 0-indexed integer array time where time[i] denotes
how many months it takes to complete the (i+1)th course.
You must find the minimum number of months needed to complete all the courses following these rules:
You may start taking a course at any time if the prerequisites are met.
Any number of courses can be taken at the same time.
Return the minimum number of months needed to complete all the courses.
Note: The test cases are generated such that it is possible to complete every course
(i.e., the graph is a directed acyclic graph).
*/

func minimumTime(n int, relations [][]int, time []int) int {
	var res int
	// graph[i] represents the courses the i'th course needs
	var graph = make([][]int, n)
	// fTime[i] represesnts the minimal time needed to complete the i'th course
	var fTime = make([]int, n)
	// solve(i) recursively finds the minimal time to complete the i'th course
	var solve func(int) int

	// Building the dependency graph
	for _, r := range relations {
		graph[r[1]-1] = append(graph[r[1]-1], r[0]-1)
	}

	solve = func(cur int) int {
		var res int
		// If we already found the minimal time, return it
		if fTime[cur] != 0 {
			return fTime[cur]
		}

		// Loop over all the course's dependencies
		for _, d := range graph[cur] {
			// Find out their minimal completion time
			dTime := solve(d)
			// Pick the maximum of those times, since you need to finish all of them
			if dTime > res {
				res = dTime
			}
		}

		// Write to cache
		fTime[cur] = res + time[cur]
		return fTime[cur]
	}

	// Loop over all the courses
	for i := range fTime {
		// Find the minimal time for each
		solve(i)
		// Pick the maximum time (the last course that needs to be completed)
		if fTime[i] > res {
			res = fTime[i]
		}
	}

	return res
}
