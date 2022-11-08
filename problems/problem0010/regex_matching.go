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
		if pattern[i] == '*' {
			// For a * we point the cell to itself
			cur.Routes = append(cur.Routes, Route{pattern[i-1], RealRoute, cur})
			cur = &RegeCell{Special: cur.Special, Routes: []Route{{'.', FakeRoute, cur}}}
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
			if rt.Match == str[0] || rt.Match == '.' {
				if rt.Dest.MatchRegex(str[1:]) {
					return true
				}
			}
		} else {
			//Fake Routes are like regular routes
			//but they don't decrement the string
			if rt.Match == str[0] || rt.Match == '.' {
				if rt.Dest.MatchRegex(str) {
					return true
				}
			}
		}
	}
	return false
}
