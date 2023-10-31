package problem2433

/*
You are given an integer array pref of size n. Find and return the array arr of size n that satisfies:
pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i].
Note that ^ denotes the bitwise-xor operation.
It can be proven that the answer is unique.
*/

func findArray(pref []int) []int {
	var prev int = pref[0]
	for i := 1; i < len(pref); i++ {
		pref[i] = prev ^ pref[i]
		prev = prev ^ pref[i]
	}
	return pref
}
