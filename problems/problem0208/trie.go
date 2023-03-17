package problem0208

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings.
There are various applications of this data structure, such as autocomplete and spellchecker.
Implement the Trie class:
    Trie() Initializes the trie object.
    void insert(String word) Inserts the string word into the trie.
    boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
    boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/

type Trie struct {
	Next   [26]*Trie
	IsWord bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	var head = t
	for i := range word {
		if head.Next[word[i]-'a'] != nil {
			head = head.Next[word[i]-'a']
		} else {
			head.Next[word[i]-'a'] = &Trie{}
			head = head.Next[word[i]-'a']
		}
	}
	head.IsWord = true
}

func (t *Trie) Search(word string) bool {
	var head = t
	for i := range word {
		if head.Next[word[i]-'a'] != nil {
			head = head.Next[word[i]-'a']
		} else {
			return false
		}
	}
	return head.IsWord
}

func (t *Trie) StartsWith(prefix string) bool {
	var head = t
	for i := range prefix {
		if head.Next[prefix[i]-'a'] != nil {
			head = head.Next[prefix[i]-'a']
		} else {
			return false
		}
	}
	return true
}
