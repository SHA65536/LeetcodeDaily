package problem0010

/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).
*/

const (
	RegularCell = iota
	EndCell
)

const (
	RealRoute = iota
	FakeRoute
)

const (
	starByte = byte(42) // Byte representation of '*'
	dotByte  = byte(46) // Byte representation of '.'
)

func isMatch(s string, p string) bool {
	return MakeRegex(p).MatchRegex(s)
}

type RegeCell struct {
	Special int
	Routes  []Route
}

type Route struct {
	Match byte
	Type  int
	Dest  *RegeCell
}

// MakeRegex makes the NFA from a pattern
func MakeRegex(pattern string) *RegeCell {
	var cur *RegeCell = &RegeCell{Special: EndCell}
	for i := len(pattern) - 1; i >= 0; i-- {
		if pattern[i] == starByte {
			// For a * we point the cell to itself
			cur.Routes = append(cur.Routes, Route{pattern[i-1], RealRoute, cur})
			if cur.Special == EndCell {
				cur = &RegeCell{Special: EndCell, Routes: []Route{{dotByte, FakeRoute, cur}}}
			} else {
				cur = &RegeCell{Routes: []Route{{dotByte, FakeRoute, cur}}}
			}
			i--
		} else {
			// For a regular character or wildcard we make a new cell
			cur = &RegeCell{Routes: []Route{{pattern[i], RealRoute, cur}}}
		}
	}
	return cur
}

func (rc *RegeCell) MatchRegex(str string) bool {
	// Exiting if string is empty
	if str == "" {
		return rc.Special == EndCell
	}
	// Because * can make multiple routes,
	// we need to check every one
	for _, rt := range rc.Routes {
		if rt.Type == RealRoute {
			if rt.Match == str[0] || rt.Match == dotByte {
				if rt.Dest.MatchRegex(str[1:]) {
					return true
				}
			}
		} else {
			if rt.Match == str[0] || rt.Match == dotByte {
				if rt.Dest.MatchRegex(str) {
					return true
				}
			}
		}
	}
	return false
}
