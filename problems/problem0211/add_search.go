package problem0211

/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.
Implement the WordDictionary class:
    WordDictionary() Initializes the object.
    void addWord(word) Adds word to the data structure, it can be matched later.
    bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise.
	word may contain dots '.' where dots can be matched with any letter.

*/

type WordDictionary struct {
	Next   [26]*WordDictionary
	IsWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (w *WordDictionary) AddWord(word string) {
	var head = w
	for i := range word {
		if head.Next[word[i]-'a'] != nil {
			head = head.Next[word[i]-'a']
		} else {
			head.Next[word[i]-'a'] = &WordDictionary{}
			head = head.Next[word[i]-'a']
		}
	}
	head.IsWord = true
}

func (w *WordDictionary) Search(word string) bool {
	var searchHelper func(cur *WordDictionary, pos int) bool

	searchHelper = func(cur *WordDictionary, pos int) bool {
		if pos == len(word) {
			return cur.IsWord
		}
		if word[pos] != '.' {
			if n := cur.Next[word[pos]-'a']; n != nil {
				return searchHelper(n, pos+1)
			}
		} else {
			for _, n := range cur.Next {
				if n != nil && searchHelper(n, pos+1) {
					return true
				}
			}
		}
		return false
	}

	return searchHelper(w, 0)
}
