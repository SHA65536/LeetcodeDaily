package problem2300

import "sort"

/*
You are given two positive integer arrays spells and potions, of length n and m respectively,
where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.
You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.
Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.
*/

func successfulPairs(spells []int, potions []int, success int64) []int {
	var res = make([]int, len(spells))
	// spellPos remembers the position of a spell based on a potion
	var spellPos = map[int][]int{}
	for i, s := range spells {
		spellPos[s] = append(spellPos[s], i)
	}
	// Sort spells and potions increasing
	sort.Ints(spells)
	sort.Ints(potions)
	var lastPot int
	// Start from the strongest spell
	for i := len(spells) - 1; i >= 0; i-- {
		if i > 0 && spells[i] == spells[i-1] {
			// If we found a duplicate, ignore
			continue
		}
		// Start from the weakest potion
		for j := lastPot; j < len(potions); j++ {
			lastPot = j
			// If succesful, it will be succesful with all stronger potions
			if int64(potions[lastPot])*int64(spells[i]) >= success {
				// Update result with original index
				for _, r := range spellPos[spells[i]] {
					res[r] = len(potions) - lastPot
				}
				break
			}
		}
	}
	return res
}

func successfulPairsTwoSum(spells []int, potions []int, success int64) []int {
	var res = make([]int, len(spells))
	var oldSpells = make([]int, len(spells))
	var count = map[int]int{}

	copy(oldSpells, spells)
	sort.Ints(potions)
	sort.Ints(spells)

	var n, m, j = len(spells), len(potions), len(potions) - 1
	for i := 0; i < n; i++ {
		for j >= 0 && int64(spells[i])*int64(potions[j]) >= success {
			j--
		}
		count[spells[i]] = m - j - 1
	}
	for i := 0; i < n; i++ {
		res[i] = count[oldSpells[i]]
	}
	return res
}
