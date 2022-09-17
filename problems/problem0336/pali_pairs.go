package problem0336

/*
Given a list of unique words
Return all the pairs of the distinct indices (i, j) in the given list,
so that the concatenation of the two words words[i] + words[j] is a palindrome.
*/

type TrieNode struct {
	Next  []*TrieNode // Next letters from current one
	Index int         // Index of word in words (-1 if not a word)
	List  []int       //
}

// New empty TrieNode
func MakeTrie() *TrieNode {
	return &TrieNode{
		Next:  make([]*TrieNode, 26),
		Index: -1,
		List:  []int{},
	}
}

func addWord(root *TrieNode, word string, index int) {
	for i := len(word) - 1; i >= 0; i-- {
		j := word[i] - 'a'
		if root.Next[j] == nil {
			root.Next[j] = MakeTrie()
		}
		if isPali(word, 0, i) {
			root.List = append(root.List, index)
		}
		root = root.Next[j]
	}
	root.List = append(root.List, index)
	root.Index = index
}

func searchWord(words []string, i int, root *TrieNode, res *[][]int) {
	for j := range words[i] {
		if root.Index >= 0 && root.Index != i && isPali(words[i], j, len(words[i])-1) {
			*res = append(*res, []int{i, root.Index})
		}
		root = root.Next[words[i][j]-'a']
		if root == nil {
			return
		}
	}
	for _, j := range root.List {
		if i != j {
			*res = append(*res, []int{i, j})
		}
	}
}

func isPali(word string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	var res = make([][]int, 0)
	var root = MakeTrie()

	// Building
	for i := range words {
		addWord(root, words[i], i)
	}

	// Searching
	for i := range words {
		searchWord(words, i, root, &res)
	}
	return res
}
