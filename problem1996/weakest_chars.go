package problem1996

import "sort"

/*
You are playing a game that contains multiple characters, and each of the characters has two main properties: attack and defense.
You are given a 2D integer array properties where properties[i] = [attacki, defensei] represents the properties of the ith character in the game.
A character is said to be weak if any other character has both attack and defense levels strictly greater than this character's attack and defense levels.
More formally, a character i is said to be weak if there exists another character j where attackj > attacki and defensej > defensei.
Return the number of weak characters.
*/

func numberOfWeakCharacters(properties [][]int) int {
	var result, max int
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})
	for i := range properties {
		if max > properties[i][1] {
			result++
		} else {
			max = properties[i][1]
		}
	}
	return result
}
