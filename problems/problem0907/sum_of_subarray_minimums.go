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

func sumSubarrayMins(arr []int) int {
	var res int
	var inDbStackP, inDbStackN = &DbStack{}, &DbStack{}
	var left, right = make([]int, len(arr)), make([]int, len(arr))

	for i := range arr {
		left[i] = i + 1
		right[i] = len(arr) - i
	}

	for i := range arr {
		for !inDbStackP.Empty() && inDbStackP.Top()[0] > arr[i] {
			inDbStackP.Pop()
		}
		left[i] = i + 1
		if !inDbStackP.Empty() {
			left[i] = i - inDbStackP.Top()[1]
		}
		inDbStackP.Push([2]int{arr[i], i})

		for !inDbStackN.Empty() && inDbStackN.Top()[0] > arr[i] {
			x := inDbStackN.Pop()
			right[x[1]] = i - x[1]
		}
		inDbStackN.Push([2]int{arr[i], i})
	}

	for i := range arr {
		res = (res + arr[i]*left[i]*right[i]) % RESMOD
	}
	return res
}

type DbStack [][2]int

func (s *DbStack) Empty() bool {
	return len(*s) == 0
}

func (s *DbStack) Top() [2]int {
	return (*s)[len(*s)-1]
}

func (s *DbStack) Pop() [2]int {
	res := s.Top()
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *DbStack) Push(v [2]int) {
	*s = append(*s, v)
}

func sumSubarrayMinsOpt(arr []int) int {
	var res, lenArr int
	var mid, leftB, rightB, count int
	lenArr = len(arr)
	st := Stack{}
	for i := 0; i <= lenArr; i++ {
		for !st.Empty() && (i == lenArr || arr[st.Top()] >= arr[i]) {
			mid = st.Pop()
			leftB, rightB = -1, i
			if len(st) > 0 {
				leftB = st.Top()
			}
			count = (rightB - mid) * (mid - leftB) % RESMOD
			res = (res + count*arr[mid]) % RESMOD
		}
		st.Push(i)
	}
	return res
}

type Stack []int

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	res := s.Top()
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}
