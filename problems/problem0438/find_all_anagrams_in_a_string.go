package problem0438

/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

func findAnagrams(s string, p string) []int {
	var results = []int{}
	var wanted, cur [26]byte
	if len(s) < len(p) {
		return results
	}
	// Finding desired word frequency and frequency of first window
	for i := range p {
		wanted[p[i]-'a']++
		cur[s[i]-'a']++
	}
	// Checking if first window matches
	if wanted == cur {
		results = append(results, 0)
	}
	// Moving the window
	for i := 0; i < len(s)-len(p); i++ {
		cur[s[i]-'a']--        // Decrementing the letter on the left
		cur[s[i+len(p)]-'a']++ // Incrementing the letter on the right
		if wanted == cur {     // Checking frequency
			results = append(results, i+1)
		}
	}
	return results
}

func findAnagramsHash(s string, p string) []int {
	var res = []int{}
	var freq = map[byte]int{}
	var i, j int
	for i := range p {
		freq[p[i]]++
	}

	for j < len(s) {
		chCnt, ok := freq[s[j]]
		if !ok {
			for i < j {
				freq[s[i]]++
				i++
			}
			i++
			j++
		} else if chCnt == 0 {
			freq[s[i]]++
			i++
		} else {
			freq[s[j]]--
			if freq[s[j]] == 0 && (j-i+1) == len(p) {
				res = append(res, i)
			}
			j++
		}
	}
	return res
}
