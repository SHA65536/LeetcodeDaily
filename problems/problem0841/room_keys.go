package problem0841

/*
There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0.
Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.
When you visit a room, you may find a set of distinct keys in it.
Each key has a number on it, denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.
Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i,
return true if you can visit all the rooms, or false otherwise.
*/

func canVisitAllRooms(rooms [][]int) bool {
	var visited = map[int]bool{0: true}
	var keyStack = Stack{}
	keyStack.Push(rooms[0]...)
	for keyStack.Len() > 0 && len(rooms) != len(visited) {
		if v := keyStack.Pop(); !visited[v] {
			visited[v] = true
			for _, k := range rooms[v] {
				if !visited[k] {
					keyStack.Push(k)
				}
			}
		}
	}
	return len(rooms) == len(visited)
}

type Stack []int

func (s *Stack) Pop() int {
	r := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return r
}

func (s *Stack) Push(x ...int) {
	*s = append(*s, x...)
}

func (s *Stack) Len() int {
	return len(*s)
}
