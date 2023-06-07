package problem1318

/*
Given 3 positives numbers a, b and c.
Return the minimum flips required in some bits of a and b to make ( a OR b == c ). (bitwise OR operation).
Flip operation consists of change any single bit 1 to 0 or change the bit 0 to 1 in their binary representation.
*/

func minFlips(a, b, c int) int {
	var res int
	var pos, i uint32
	var A, B, C = uint32(a), uint32(b), uint32(c)
	for i = 0; i < 32; i++ {
		pos = 1 << i
		if C&pos == 0 {
			// Target has 0 at position i

			// Both A and B should have 0 in position i
			if A&pos != 0 {
				res++
			}
			if B&pos != 0 {
				res++
			}
		} else {
			// Target has 1 at position i

			// At least one of A or B should have 1
			// in position i
			if A&pos == 0 && B&pos == 0 {
				res++
			}
		}
	}
	return res
}
