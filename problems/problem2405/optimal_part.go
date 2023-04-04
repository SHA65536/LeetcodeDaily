package problem2405

/*
Given a string s, partition the string into one or more substrings such that the characters in each substring are unique.
That is, no letter appears in a single substring more than once.
Return the minimum number of substrings in such a partition.
Note that each character should belong to exactly one substring in a partition.
*/

func partitionStringMap(s string) int {
	var res int
	var freq = map[byte]bool{}

	for i := range s {
		if freq[s[i]] {
			res++
			freq = map[byte]bool{}
		}
		freq[s[i]] = true
	}

	return res + 1
}

func partitionStringArr(s string) int {
	var res int
	var freq [26]bool

	for i := range s {
		if freq[s[i]-'a'] {
			res++
			freq = [26]bool{}
		}
		freq[s[i]-'a'] = true
	}

	return res + 1
}

func partitionStringBit(s string) int {
	var res int
	var freq uint32

	for i := range s {
		if freq&(1<<(s[i]-'a')) > 0 {
			res++
			freq = 0
		}
		freq |= 1 << (s[i] - 'a')
	}

	return res + 1
}
