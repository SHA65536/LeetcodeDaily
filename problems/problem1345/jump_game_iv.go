package problem1345

/*
Given an array of integers arr, you are initially positioned at the first index of the array.
In one step you can jump from index i to index:
    i + 1 where: i + 1 < arr.length.
    i - 1 where: i - 1 >= 0.
    j where: arr[i] == arr[j] and i != j.
Return the minimum number of steps to reach the last index of the array.
Notice that you can not jump outside of the array at any time.
*/

func minJumps(arr []int) int {
	var res int
	if len(arr) == 1 {
		return 0
	}
	// valJump[i] represents the indexes you can jump to from i
	var valJumps = map[int][]int{}
	for i := range arr {
		valJumps[arr[i]] = append(valJumps[arr[i]], i)
	}
	// seen[i] represents if we've visited the i element
	var seen = make([]bool, len(arr))
	seen[0] = true

	// cur and nxt are the steps of the bfs
	var cur, nxt []int
	cur = []int{0}

	for len(cur) > 0 {
		res++
		nxt = []int{}
		for _, n := range cur {
			// Jump forward
			if n+1 == len(arr)-1 {
				return res
			} else if !seen[n+1] {
				seen[n+1] = true
				nxt = append(nxt, n+1)
			}
			// Jump backwards
			if n-1 >= 0 && !seen[n-1] {
				seen[n-1] = true
				nxt = append(nxt, n-1)
			}
			// Jump by value
			valArr := valJumps[arr[n]] // Temp variable to have less map access
			// Looping from end to start incase the end of the array is in there
			for ni := len(valArr) - 1; ni >= 0; ni-- {
				if valArr[ni] == len(arr)-1 {
					return res
				} else if !seen[valArr[ni]] {
					nxt = append(nxt, valArr[ni])
					seen[valArr[ni]] = true
				}
			}
			// Delete this so we don't need to loop over it again
			delete(valJumps, arr[n])
		}
		cur = nxt
	}
	return res
}
