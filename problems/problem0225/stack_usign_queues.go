package problem0225

/*
Implement a last-in-first-out (LIFO) stack using only two queues.
The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).
Implement the MyStack class:
    void push(int x) Pushes element x to the top of the stack.
    int pop() Removes the element on the top of the stack and returns it.
    int top() Returns the element on the top of the stack.
    boolean empty() Returns true if the stack is empty, false otherwise.
Notes:
    You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
    Depending on your language, the queue may not be supported natively.
	You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.
*/

type MyStack struct {
	Q *MyQueue
}

func Constructor() MyStack {
	return MyStack{&MyQueue{}}
}

func (s *MyStack) Push(x int) {
	s.Q.Push(x)
	// Reverse the queue
	for i := 0; i < s.Q.Len()-1; i++ {
		s.Q.Push(s.Q.Pop())
	}
}

func (s *MyStack) Pop() int {
	return s.Q.Pop()
}

func (s *MyStack) Top() int {
	return s.Q.Peek()
}

func (s *MyStack) Empty() bool {
	return s.Q.Empty()
}

type MyQueue struct {
	data []int
}

func (q *MyQueue) Len() int    { return len(q.data) }
func (q *MyQueue) Empty() bool { return len(q.data) == 0 }
func (q *MyQueue) Push(x int)  { q.data = append(q.data, x) }
func (q *MyQueue) Peek() int   { return q.data[0] }
func (q *MyQueue) Pop() int {
	var res = q.data[0]
	q.data = q.data[1:]
	return res
}
