package problem0017

/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.
*/

var keysMap = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var res = []string{}
	if digits == "" {
		return []string{}
	}
	combos := letterCombinations(digits[1:])
	if len(combos) == 0 {
		combos = []string{""}
	}
	for _, letter := range keysMap[rune(digits[0])] {
		cur := string(letter)
		for _, combo := range combos {
			res = append(res, cur+combo)
		}
	}
	return res
}
