package problem1921

import "sort"

/*
You are playing a video game where you are defending your city from a group of n monsters.
You are given a 0-indexed integer array dist of size n, where dist[i] is the initial distance in kilometers of the ith monster from the city.
The monsters walk toward the city at a constant speed.
The speed of each monster is given to you in an integer array speed of size n, where speed[i] is the speed of the ith monster in kilometers per minute.
You have a weapon that, once fully charged, can eliminate a single monster.
However, the weapon takes one minute to charge. The weapon is fully charged at the very start.
You lose when any monster reaches your city.
If a monster reaches the city at the exact moment the weapon is fully charged, it counts as a loss, and the game ends before you can use your weapon.
Return the maximum number of monsters that you can eliminate before you lose, or n if you can eliminate all the monsters before they reach the city.
*/

func eliminateMaximum(dist []int, speed []int) int {
	// Calculate time of arrival
	for i := range dist {
		if dist[i]%speed[i] > 0 {
			dist[i] = 1 + (dist[i] / speed[i])
		} else {
			dist[i] = dist[i] / speed[i]
		}
	}

	// Sort by closest arrival
	sort.Ints(dist)

	// Check if we can keep up with eliminating
	var turn int
	for i := range dist {
		if turn >= dist[i] {
			return turn
		}
		turn++
	}

	return turn
}

func eliminateMaximumNoSort(dist []int, speed []int) int {
	// Calculate time of arrival
	for i := range dist {
		if dist[i]%speed[i] > 0 {
			dist[i] = 1 + (dist[i] / speed[i])
		} else {
			dist[i] = dist[i] / speed[i]
		}
		speed[i] = 0 // Reset for further use
	}

	// Calculate how many arrive each minute
	for _, t := range dist {
		if t < len(dist) { // If the arrive after n minutes we can eliminate all
			speed[t]++
		}
	}

	// Check if we can keep up with the monsters
	for i := 1; i < len(speed); i++ {
		speed[i] += speed[i-1]
		if speed[i] > i {
			return i
		}
	}

	return len(speed)
}
