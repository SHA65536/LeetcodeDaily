package problem1376

/*
A company has n employees with a unique ID for each employee from 0 to n - 1.
The head of the company is the one with headID.
Each employee has one direct manager given in the manager array where manager[i] is the direct manager of the i-th employee, manager[headID] = -1.
Also, it is guaranteed that the subordination relationships have a tree structure.
The head of the company wants to inform all the company employees of an urgent piece of news.
He will inform his direct subordinates, and they will inform their subordinates, and so on until all employees know about the urgent news.
The i-th employee needs informTime[i] minutes to inform all of his direct subordinates
(i.e., After informTime[i] minutes, all his direct subordinates can start spreading the news).
Return the number of minutes needed to inform all the employees about the urgent news.
*/

func numOfMinutes(n, headID int, manager, informTime []int) int {
	// graph[i] is a list of employees that
	// the ith employee is a manager of
	var graph = make([][]int, n)

	// Building manager graph
	for i := range manager {
		if manager[i] != -1 {
			graph[manager[i]] = append(graph[manager[i]], i)
		}
	}

	// Getting notify time of head
	return notifyTime(headID, graph, informTime)
}

// notifyTime returns the time the nth employ takes to inform
// all of it's subordinates
func notifyTime(n int, graph [][]int, informTime []int) int {
	var res int = 0

	// Take the maximum time
	for i := range graph[n] {
		res = max(res, notifyTime(graph[n][i], graph, informTime))
	}

	// If no subordinate time, it doesn't have to spend
	// time notifying other employees
	if len(graph[n]) == 0 {
		return res
	}

	return res + informTime[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
