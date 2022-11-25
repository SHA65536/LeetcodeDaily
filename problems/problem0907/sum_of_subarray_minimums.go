package problem0907

/*
Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr.
Since the answer may be large, return the answer modulo 10^9 + 7.
*/

const RESMOD = 1000000007
const MAXVAL = 30001

func sumSubarrayMinsNaive(arr []int) int {
	var res, start, end int
	var curMin int
	for start = 0; start < len(arr); start++ { // Start index of subarray
		curMin = MAXVAL
		for end = start; end < len(arr); end++ { // End index of subarray
			if arr[end] < curMin { // Discovering a new min
				curMin = arr[end]
			}
			res = (res + curMin) % RESMOD
		}
	}
	return res
}

type Stack [][2]int

func sumSubarrayMins(arr []int) int {
	var res int
	var inStackP, inStackN = &Stack{}, &Stack{}
	var left, right = make([]int, len(arr)), make([]int, len(arr))

	for i := range arr {
		left[i] = i + 1
		right[i] = len(arr) - i
	}

	for i := range arr {
		for !inStackP.Empty() && inStackP.Top()[0] > arr[i] {
			inStackP.Pop()
		}
		left[i] = i + 1
		if !inStackP.Empty() {
			left[i] = i - inStackP.Top()[1]
		}
		inStackP.Push([2]int{arr[i], i})

		for !inStackN.Empty() && inStackN.Top()[0] > arr[i] {
			x := inStackN.Pop()
			right[x[1]] = i - x[1]
		}
		inStackN.Push([2]int{arr[i], i})
	}

	for i := range arr {
		res = (res + arr[i]*left[i]*right[i]) % RESMOD
	}
	return res
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Top() [2]int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() [2]int {
	res := s.Top()
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Push(v [2]int) {
	*s = append(*s, v)
}
