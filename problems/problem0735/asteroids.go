package problem0735

/*
We are given an array asteroids of integers representing asteroids in a row.
For each asteroid, the absolute value represents its size,
and the sign represents its direction (positive meaning right, negative meaning left).
Each asteroid moves at the same speed.
Find out the state of the asteroids after all collisions.
If two asteroids meet, the smaller one will explode.
If both are the same size, both will explode.
Two asteroids moving in the same direction will never meet.
*/

func asteroidCollision(asteroids []int) []int {
	var stack = Stack[int]{}
	for i := 0; i < len(asteroids); i++ {
		cur := asteroids[i]
		// If the stack is empty, just add the asteroid
		// If the last element is moving left, just add the asteroid
		// If the direction is the same, just add the asteroid
		if len(stack) == 0 || stack.Peek() < 0 || stack.Peek() > 0 && cur > 0 {
			stack.Push(cur)
			continue
		}

		if abs(stack.Peek()) == abs(cur) {
			// If asteroids are the same size, remove both
			stack.Pop()
		} else if abs(stack.Peek()) < abs(cur) {
			// If the left asteroid is smaller, remove it and re-evaluate
			stack.Pop()
			i--
		}
		// If the left asteroid is bigger, we don't need to do anything
	}

	return stack
}

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() T {
	var res = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
