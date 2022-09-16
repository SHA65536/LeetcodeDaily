package problem0022

/*
https://leetcode.com/problems/generate-parentheses/

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

func generateParenthesis(n int) []string {
	return genParenthesis(n, 0, 0)
}

func genParenthesis(max, open, close int) []string {
	var res = []string{}
	if open+close+1 == max*2 {
		return []string{")"}
	}
	if open == close {
		for _, str := range genParenthesis(max, open+1, close) {
			res = append(res, "("+str)
		}
	} else if open == max {
		for _, str := range genParenthesis(max, open, close+1) {
			res = append(res, ")"+str)
		}
	} else {
		for _, str := range genParenthesis(max, open+1, close) {
			res = append(res, "("+str)
		}
		for _, str := range genParenthesis(max, open, close+1) {
			res = append(res, ")"+str)
		}
	}
	return res
}
