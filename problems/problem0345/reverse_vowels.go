package problem0345

/*
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/

var vowels = map[byte]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func reverseVowels(s string) string {
	var res = []byte(s)
	for i, j := 0, len(s)-1; j > i; {
		if vowels[s[i]] && vowels[s[j]] {
			res[i], res[j] = res[j], res[i]
			i++
			j--
		}
		if !vowels[s[i]] {
			i++
		}
		if !vowels[s[j]] {
			j--
		}
	}
	return string(res)
}
