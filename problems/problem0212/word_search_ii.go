package problem0212

/*
Given an m x n board of characters and a list of strings words, return all words on the board.
Each word must be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.
*/

// Trie built from input words
type Trie struct {
	Next [26]*Trie
	Word string
}

func MakeTrie(words []string) *Trie {
	var root, cur *Trie
	var letter byte
	root = &Trie{}
	for i := range words {
		cur = root
		for j := range words[i] {
			letter = words[i][j] - 'a'
			if cur.Next[letter] == nil {
				cur.Next[letter] = &Trie{}
			}
			cur = cur.Next[letter]
		}
		cur.Word = words[i]
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	var res *[]string = &[]string{}
	var trie = MakeTrie(words)
	for i := range board {
		for j := range board[0] {
			// Checking each square with full trie
			findAt(board, i, j, trie, res)
		}
	}
	return *res
}

// findAt checks if there is any match in the trie starting at i,j
func findAt(board [][]byte, i, j int, trie *Trie, res *[]string) {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
		// Check if outside the board
		return
	}
	var curLet = board[i][j]
	var curTrie *Trie
	if curLet == 0 || trie.Next[curLet-'a'] == nil {
		// Check if cell is already visited, or if no word matches current char
		return
	}
	// Current trie node matching current travel
	curTrie = trie.Next[curLet-'a']
	if curTrie.Word != "" {
		// If we found a match, add to res, we keep
		// going because there might be another word
		*res = append(*res, curTrie.Word)
		curTrie.Word = ""
	}

	// Marking current cell as visisted
	board[i][j] = 0
	// Checking adjacent squares
	findAt(board, i+1, j, curTrie, res)
	findAt(board, i, j+1, curTrie, res)
	findAt(board, i-1, j, curTrie, res)
	findAt(board, i, j-1, curTrie, res)
	// Unmarking current cell for future travels
	board[i][j] = curLet
}
