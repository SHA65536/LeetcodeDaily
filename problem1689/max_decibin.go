package problem1689

/*
https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers

A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros.
For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.
*/

func minPartitions(n string) int {
	// Since all deci-binary numbers only have the digits 1 to 0
	// we just need to look for the biggest digit
	var maxbyte byte = 48
	for idx := range n {
		if maxbyte == 57 {
			return 9
		}
		if n[idx] > maxbyte {
			maxbyte = n[idx]
		}
	}
	return (int(maxbyte) - 48)
}
