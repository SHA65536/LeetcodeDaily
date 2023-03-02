package problem0443

/*
Given an array of characters chars, compress it using the following algorithm:
Begin with an empty string s. For each group of consecutive repeating characters in chars:
    If the group's length is 1, append the character to s.
    Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
Note that group lengths that are 10 or longer will be split into multiple characters in chars.
After you are done modifying the input array, return the new length of the array.
You must write an algorithm that uses only constant extra space.
*/

func compress(chars []byte) int {
	var write, read, cons int
	var prev byte = chars[0]
	cons = 1
	for read = 1; read < len(chars); read++ {
		if chars[read] == prev {
			cons++
		} else {
			chars[write] = prev
			write++
			if cons > 1 {
				temp := intToBytes(cons)
				for i := len(temp) - 1; i >= 0; i-- {
					chars[write] = temp[i]
					write++
				}
			}
			prev = chars[read]
			cons = 1
		}
	}
	chars[write] = prev
	write++
	if cons > 1 {
		temp := intToBytes(cons)
		for i := len(temp) - 1; i >= 0; i-- {
			chars[write] = temp[i]
			write++
		}
	}
	return write
}

func intToBytes(in int) []byte {
	var res []byte
	for in > 0 {
		res = append(res, byte((in%10)+'0'))
		in /= 10
	}
	return res
}
