package problem1834

import (
	"container/heap"
	"sort"
)

/*
You are given n​​​​​​ tasks labeled from 0 to n - 1 represented by a 2D integer array tasks,
where tasks[i] = [enqueueTimei, processingTimei] means that the i​​​​​​th​​​​ task will be available to process at enqueueTime i
and will take processingTime i to finish processing.
You have a single-threaded CPU that can process at most one task at a time and will act in the following way:
If the CPU is idle and there are no available tasks to process, the CPU remains idle.
If the CPU is idle and there are available tasks, the CPU will choose the one with the shortest processing time.
If multiple tasks have the same shortest processing time, it will choose the task with the smallest index.
Once a task is started, the CPU will process the entire task without stopping.
The CPU can finish a task then start a new one instantly.
Return the order in which the CPU will process the tasks
*/

type Task struct {
	Enqueue int
	Process int
	Idx     int
}

func getOrder(tasks [][]int) []int {
	var res = make([]int, 0, len(tasks))
	var taskStack = MakeHeap(MinHeap)
	var newTasks = make([]Task, len(tasks))
	for i := range tasks {
		newTasks[i] = Task{
			Enqueue: tasks[i][0],
			Process: tasks[i][1],
			Idx:     i,
		}
	}
	// Sorting by enqueue time
	sort.SliceStable(newTasks, func(i, j int) bool { return newTasks[i].Enqueue < newTasks[j].Enqueue })

	// Adding to res
	var cur, time int
	for time = newTasks[0].Enqueue; cur < len(tasks); {
		// Adding all tasks that can be done right now
		for cur < len(newTasks) && newTasks[cur].Enqueue <= time {
			heap.Push(taskStack, newTasks[cur])
			cur++
		}
		// If no tasks can be done
		if taskStack.Len() == 0 {
			// Set the time for the next task
			time = newTasks[cur].Enqueue
		} else {
			// Adding the next task
			v := heap.Pop(taskStack).(Task)
			res = append(res, v.Idx)
			time += v.Process
		}

	}

	// Adding remaining tasks
	for taskStack.Len() > 0 {
		v := heap.Pop(taskStack).(Task)
		res = append(res, v.Idx)
	}
	return res
}

type Heap struct {
	Values   []Task
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool {
	if h.Values[i].Process == h.Values[j].Process {
		return h.LessFunc(h.Values[i].Idx, h.Values[j].Idx)
	}
	return h.LessFunc(h.Values[i].Process, h.Values[j].Process)
}
func (h *Heap) Swap(i, j int) { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int      { return len(h.Values) }
func (h *Heap) Peek() Task    { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(Task)) }

func MakeHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

func MaxHeap(i, j int) bool { return i > j }
func MinHeap(i, j int) bool { return i < j }
