package problem1964

import "sort"

/*
You want to build some obstacle courses.
You are given a 0-indexed integer array obstacles of length n, where obstacles[i] describes the height of the ith obstacle.
For every index i between 0 and n - 1 (inclusive), find the length of the longest obstacle course in obstacles such that:
    You choose any number of obstacles between 0 and i inclusive.
    You must include the ith obstacle in the course.
    You must put the chosen obstacles in the same order as they appear in obstacles.
    Every obstacle (except the first) is taller than or the same height as the obstacle immediately before it.
Return an array ans of length n, where ans[i] is the length of the longest obstacle course for index i as described above.
*/

type stack []int

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var s stack
	for i := range obstacles {
		var x = obstacles[i]
		if s.empty() || s.peek() <= x {
			s.push(x)
			obstacles[i] = len(s)
		} else {
			idx := sort.Search(len(s), func(j int) bool { return (s)[j] > x })
			s[idx] = x
			obstacles[i] = idx + 1
		}
	}
	return obstacles
}

func (s *stack) push(x int)  { *s = append(*s, x) }
func (s *stack) peek() int   { return (*s)[len(*s)-1] }
func (s *stack) empty() bool { return len(*s) == 0 }
