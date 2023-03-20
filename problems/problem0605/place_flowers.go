package problem0605

/*
You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.
Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.
*/

func canPlaceFlowers(f []int, n int) bool {
	// Looping until we run out of flowers, or run out of space
	for i := 0; n > 0 && i < len(f); i++ {
		// f[i] == 0: Current spot is not a flower
		// (i == 0 || f[i-1] == 0): There is no left adjacent, or there is no flower on the left
		// i == len(f)-1 || f[i+1] == 0): There is no right adjacent, or there is no flower on the right
		if f[i] == 0 && (i == 0 || f[i-1] == 0) && (i == len(f)-1 || f[i+1] == 0) {
			f[i] = 1
			n--
		}
	}
	// Return true if we planted all flowers
	return n <= 0
}
