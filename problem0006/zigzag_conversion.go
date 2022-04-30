package problem0006

/*
https://leetcode.com/problems/zigzag-conversion/

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows
*/

func convert(s string, numRows int) string {
	var res string
	var layidx, inc = 0, -1
	var layers = make([][]byte, numRows)
	if numRows == 1 {
		return s
	}
	for idx := range s {
		layers[layidx] = append(layers[layidx], s[idx])
		if layidx == numRows-1 || layidx == 0 {
			inc *= -1
		}
		layidx += inc
	}
	for _, layer := range layers {
		res += string(layer)
	}
	return res
}
