package problem2306

/*
You are given an array of strings ideas that represents a list of names to be used in the process of naming a company.
The process of naming a company is as follows:
    Choose 2 distinct names from ideas, call them ideaA and ideaB.
    Swap the first letters of ideaA and ideaB with each other.
    If both of the new names are not found in the original ideas, then the name ideaA ideaB
	(the concatenation of ideaA and ideaB, separated by a space) is a valid company name.
    Otherwise, it is not a valid name.
Return the number of distinct valid names for the company.
*/

func distinctNames(ideas []string) int64 {
	var res int64
	// prefixes[s] represents the set of strings that start with s
	var prefixes = [26]map[string]bool{}
	for _, s := range ideas {
		// Populating the prefixes array
		if prefixes[s[0]-'a'] != nil {
			prefixes[s[0]-'a'][s[1:]] = true
		} else {
			prefixes[s[0]-'a'] = map[string]bool{s[1:]: true}
		}
	}

	// For each prefix group
	for i := 0; i < len(prefixes); i++ {
		// Compare against other prefix groups that haven't been compared before
		for j := i + 1; j < len(prefixes); j++ {
			// Find duplicate suffixes of both groups
			dup := duplicates(prefixes[i], prefixes[j])
			// Multiply the number of unique suffixes between the groups
			res += int64(2 * (len(prefixes[i]) - dup) * (len(prefixes[j]) - dup))
		}
	}
	return res
}

func duplicates(a, b map[string]bool) int {
	var res int
	if len(a) > len(b) {
		a, b = b, a
	}
	for k := range a {
		if b[k] {
			res++
		}
	}
	return res
}
