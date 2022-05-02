package problem0010

import "fmt"

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).
*/

func isMatch(s string, p string) bool {
	fmt.Println(s, p, "\n==========")
	var reg = MakeRegex(p)
	for i := 0; i < len(s); i++ {
		if next, ok := reg.Routes[s[i]]; ok {
			reg = next
		} else {
			return false
		}
	}
	return reg.Special == EndCell
}

const (
	RegularCell = iota
	TrapCell
	EndCell
)

type RegeCell struct {
	Special int
	Routes  map[byte]*RegeCell
}

func MakeRegex(pattern string) *RegeCell {
	var cur *RegeCell = &RegeCell{Special: EndCell}
	fmt.Println("Making: ", pattern)
	for i := len(pattern) - 1; i >= 0; i-- {
		before := &RegeCell{Routes: map[byte]*RegeCell{pattern[i]: cur}}
		cur = before
	}
	return cur
}
