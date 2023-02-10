package problem0933

/*
You have a RecentCounter class which counts the number of recent requests within a certain time frame.
Implement the RecentCounter class:

    RecentCounter() Initializes the counter with zero recent requests.
    int ping(int t) Adds a new request at time t, where t represents some time in milliseconds,
	and returns the number of requests that has happened in the past 3000 milliseconds (including the new request).
	Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].

It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.
*/

type RecentCounter struct {
	Queue *CircularQueue
}

func Constructor() RecentCounter {
	q := MakeQueue(3001)
	return RecentCounter{
		Queue: &q,
	}
}

func (rp *RecentCounter) Ping(t int) int {
	last := t - 3000
	for !rp.Queue.IsEmpty() && rp.Queue.Front() < last {
		rp.Queue.DeQueue()
	}
	rp.Queue.EnQueue(t)
	return rp.Queue.Count
}

// Circular Queue from problem 622
type CircularQueue struct {
	Queue    []int // Actual slice with data
	Position int   // Position of first element
	Count    int   // Number of elements in queue
}

func MakeQueue(k int) CircularQueue {
	return CircularQueue{
		Queue:    make([]int, k),
		Position: 0,
		Count:    0,
	}
}

func (q *CircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.Queue[(q.Position+q.Count)%len(q.Queue)] = value
	q.Count++
	return true
}

func (q *CircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.Position = (q.Position + 1) % len(q.Queue)
	q.Count--
	return true
}

func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Queue[q.Position]
}

func (q *CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Queue[(q.Position+q.Count-1)%len(q.Queue)]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.Count == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.Count == len(q.Queue)
}
