package problem0126

/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList,
Return all the shortest transformation sequences from beginWord to endWord, or an empty list if no such sequence exists.
Each sequence should be returned as a list of the words [beginWord, s1, s2, ..., sk].
*/

func diffChar(a, b string) int {
	var ans int
	for id := range a {
		if a[id] != b[id] {
			ans++
		}
	}
	return ans
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var ans = [][]string{}
	var mp = map[string]int{beginWord: 0}
	var queue = []string{beginWord}
	for len(queue) > 0 {
		curWord := queue[0]
		queue = queue[1:]
		for _, nextWord := range wordList {
			if _, have := mp[nextWord]; !have && diffChar(curWord, nextWord) == 1 {
				queue = append(queue, nextWord)
				mp[nextWord] = mp[curWord] + 1
			}
		}
	}

	if _, have := mp[endWord]; !have {
		return ans
	}

	searchWords(&mp, &beginWord, &ans, &wordList, endWord, []string{})

	return ans
}

func searchWords(mp *map[string]int, begin *string, ans *[][]string, wordList *[]string, curWord string, cur []string) {
	cur = append(cur, curWord)
	if (*mp)[curWord] == 1 {
		cur = append(cur, *begin)
		tmp := cur
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
		*ans = append(*ans, tmp)
		return
	}
	for _, nextWord := range *wordList {
		if _, ok := (*mp)[nextWord]; ok && diffChar(curWord, nextWord) == 1 && (*mp)[curWord] == (*mp)[nextWord]+1 {
			searchWords(mp, begin, ans, wordList, nextWord, cur)
		}
	}
}
