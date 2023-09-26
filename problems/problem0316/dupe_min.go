package problem0316

/*
Given a string s, remove duplicate letters so that every letter appears once and only once.
You must make sure your result is the smallest in lexicographical order among all possible results.
*/

func removeDuplicateLetters(s string) string {
	var last = make(map[byte]int, len(s))
	var visited = make(map[byte]bool, len(s))
	var stack = Stack(make([]byte, 0, len(s)))

	for i := range s {
		last[s[i]] = i
	}

	for i := range s {
		if visited[s[i]] {
			continue
		}
		for len(stack) > 0 && stack.Peek() > s[i] && last[stack.Peek()] > i {
			delete(visited, stack.Pop())
		}
		stack.Push(s[i])
		visited[s[i]] = true
	}

	return string(stack)
}

func removeDuplicateLettersArr(s string) string {
	var last [26]int
	var visited [26]bool
	var stack = Stack(make([]byte, 0, len(s)))

	for i := range s {
		last[s[i]-'a'] = i + 1
	}

	for i := range s {
		if visited[s[i]-'a'] {
			continue
		}
		for len(stack) > 0 && stack.Peek() > s[i] && last[stack.Peek()-'a']-1 > i {
			visited[stack.Pop()-'a'] = false
		}
		stack.Push(s[i])
		visited[s[i]-'a'] = true
	}

	return string(stack)
}

type Stack []byte

func (s *Stack) Push(a byte) {
	*s = append(*s, a)
}

func (s *Stack) Pop() byte {
	t := s.Peek()
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *Stack) Peek() byte {
	return (*s)[len(*s)-1]
}
