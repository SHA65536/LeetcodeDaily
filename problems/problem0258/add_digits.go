package problem0258

/*
Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
*/

func addDigits(num int) int {
	for num > 9 {
		var nxt int
		for num > 0 {
			nxt += num % 10
			num /= 10
		}
		num = nxt
	}
	return num
}

func addDigitsMod(num int) int {
	if num == 0 {
		return 0
	} else if num%9 == 0 {
		return 9
	}
	return num % 9
}
