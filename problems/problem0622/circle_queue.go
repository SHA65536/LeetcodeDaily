package problem0622

/*
Design your implementation of the circular queue.
The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle
and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue.
In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue.
But using the circular queue, we can use the space to store new values.

Implementation the MyCircularQueue class:
MyCircularQueue(k) Initializes the object with the size of the queue to be k.
int Front() Gets the front item from the queue. If the queue is empty, return -1.
int Rear() Gets the last item from the queue. If the queue is empty, return -1.
boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
boolean isEmpty() Checks whether the circular queue is empty or not.
boolean isFull() Checks whether the circular queue is full or not.
You must solve the problem without using the built-in queue data structure in your programming language.
*/

type MyCircularQueue struct {
	Queue    []int // Actual slice with data
	Position int   // Position of first element
	Count    int   // Number of elements in queue
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Queue:    make([]int, k),
		Position: 0,
		Count:    0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Queue[(this.Position+this.Count)%len(this.Queue)] = value
	this.Count++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Position = (this.Position + 1) % len(this.Queue)
	this.Count--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Position]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[(this.Position+this.Count-1)%len(this.Queue)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Count == len(this.Queue)
}
