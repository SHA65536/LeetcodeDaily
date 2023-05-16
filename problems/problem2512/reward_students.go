package problem2512

import (
	"sort"
)

/*
You are given two string arrays positive_feedback and negative_feedback,
containing the words denoting positive and negative feedback, respectively.
Note that no word is both positive and negative.
Initially every student has 0 points. Each positive word in a feedback report increases the points of a student by 3,
whereas each negative word decreases the points by 1.
You are given n feedback reports, represented by a 0-indexed string array report and a 0-indexed integer array student_id,
where student_id[i] represents the ID of the student who has received the feedback report report[i].
The ID of each student is unique.
Given an integer k, return the top k students after ranking them in non-increasing order by their points.
In case more than one student has the same points, the one with the lower ID ranks higher.
*/

type Pair struct {
	Id, Val int
}

func topStudentsMap(positive, negative, report []string, student_id []int, k int) []int {
	var feedback = make(map[string]int, len(positive)+len(negative))
	var rewards = make([]Pair, len(student_id))
	var res = make([]int, len(student_id))
	for i := range positive {
		feedback[positive[i]] = 3
	}
	for i := range negative {
		feedback[negative[i]] = -1
	}
	for i := range rewards {
		rewards[i].Id = student_id[i]
		rewards[i].Val = wordSearch(report[i], feedback)
	}
	sort.Slice(rewards, func(i, j int) bool {
		if rewards[i].Val == rewards[j].Val {
			return rewards[i].Id < rewards[j].Id
		}
		return rewards[i].Val > rewards[j].Val
	})
	for i := range res {
		res[i] = rewards[i].Id
	}
	return res[:k]
}

func wordSearch(input string, feedback map[string]int) int {
	var res int
	var prev int
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			res += feedback[input[prev:i]]
			prev = i + 1
		}
	}
	return res + feedback[input[prev:]]
}

func topStudentsTrie(positive, negative, report []string, student_id []int, k int) []int {
	var rewards = make([]Pair, len(student_id))
	var res = make([]int, len(student_id))
	// feedback is a trie for efficient search of the words
	var feedback Trie
	// Construct the trie with positive and negative feedback words
	for i := range positive {
		feedback.Insert(positive[i], 3)
	}
	for i := range negative {
		feedback.Insert(negative[i], -1)
	}
	for i := range rewards {
		rewards[i].Id = student_id[i]
		for j := range report[i] {
			rewards[i].Val += feedback.Search(report[i], j)
		}
	}
	sort.Slice(rewards, func(i, j int) bool {
		if rewards[i].Val == rewards[j].Val {
			return rewards[i].Id < rewards[j].Id
		}
		return rewards[i].Val > rewards[j].Val
	})
	for i := range res {
		res[i] = rewards[i].Id
	}
	return res[:k]
}

type Trie struct {
	// Next[i] is the next node in the trie that starts with the
	// i'th letter of the alphabet
	Next [26]*Trie
	// Value is the value of the word (0 if not a word)
	Value int
}

// Insert updates the trie with a new word
func (t *Trie) Insert(word string, value int) {
	var head = t
	// Loop over all the letters
	for i := range word {
		if head.Next[word[i]-'a'] != nil {
			// If the next letter exists, just traverse there
			head = head.Next[word[i]-'a']
		} else {
			// If it doesn't exist, create it and traverse
			head.Next[word[i]-'a'] = &Trie{}
			head = head.Next[word[i]-'a']
		}
	}
	head.Value = value
}

func (t *Trie) Search(sentence string, pos int) int {
	var head = t
	for i := pos; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			return head.Value
		} else if head.Next[sentence[i]-'a'] == nil {
			return 0
		}
		head = head.Next[sentence[i]-'a']
	}
	return head.Value
}
