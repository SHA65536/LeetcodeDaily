package problem0133

/*
Given a reference of a node in a connected undirected graph.
Return a deep copy (clone) of the graph.
Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(start *Node) *Node {
	var seen = map[int]*Node{}
	if start == nil {
		return start
	}
	var head = &Node{Val: start.Val, Neighbors: start.Neighbors}
	seen[head.Val] = head
	var cur, nxt []*Node = []*Node{head}, nil
	for len(cur) > 0 {
		nxt = []*Node{}
		for _, n := range cur {
			var newNeighbors = []*Node{}
			for _, m := range n.Neighbors {
				if seen[m.Val] == nil {
					seen[m.Val] = &Node{Val: m.Val, Neighbors: m.Neighbors}
					nxt = append(nxt, seen[m.Val])
				}
				newNeighbors = append(newNeighbors, seen[m.Val])
			}
			n.Neighbors = newNeighbors
		}
		cur = nxt
	}
	return head
}
