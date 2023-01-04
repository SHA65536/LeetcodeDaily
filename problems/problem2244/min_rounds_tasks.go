package problem2244

/*
You are given a 0-indexed integer array tasks, where tasks[i] represents the difficulty level of a task.
In each round, you can complete either 2 or 3 tasks of the same difficulty level.
Return the minimum rounds required to complete all the tasks, or -1 if it is not possible to complete all the tasks.
*/

func minimumRounds(tasks []int) int {
	var res int
	var tmap = map[int]int{}
	for _, i := range tasks {
		tmap[i]++
	}
	for _, v := range tmap {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			res += v / 3
		} else {
			res += (v / 3) + 1
		}
	}
	return res
}
