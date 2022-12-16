package problem0232

/*
Implement a first in first out (FIFO) queue using only two stacks.
The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).
Implement the MyQueue class:
    void push(int x) Pushes element x to the back of the queue.
    int pop() Removes the element from the front of the queue and returns it.
    int peek() Returns the element at the front of the queue.
    boolean empty() Returns true if the queue is empty, false otherwise.
*/

// Stack implementation
type MyStack []int

func (s *MyStack) Push(x int) {
	*s = append(*s, x)
}

func (s *MyStack) Pop() int {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack) Len() int {
	return len(*s)
}

// Queue implementation with 2 stacks
type MyQueue struct {
	First, Second MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		First:  make(MyStack, 0),
		Second: make(MyStack, 0),
	}
}

// Push to the first
func (q *MyQueue) Push(x int) {
	q.First.Push(x)
}

// Pop from second
func (q *MyQueue) Pop() int {
	// If second is empty, transfer all the vals from first to second
	if q.Second.Len() == 0 {
		q.Switch()
	}
	return q.Second.Pop()
}

// Peek from second
func (q *MyQueue) Peek() int {
	// If second is empty, transfer all the vals from first to second
	if q.Second.Len() == 0 {
		q.Switch()
	}
	return q.Second.Peek()
}

// Both need to be checked
func (q *MyQueue) Empty() bool {
	return q.First.Len() == 0 && q.Second.Len() == 0
}

// Transfer all the values from first to second
// They will end up in reverse order
func (q *MyQueue) Switch() {
	for q.First.Len() > 0 {
		q.Second.Push(q.First.Pop())
	}
}
