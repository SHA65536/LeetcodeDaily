package problem2259

/*
You are given a string number representing a positive integer and a character digit.
Return the resulting string after removing exactly one occurrence of digit from number such that the value of the resulting string in decimal form is maximized.
The test cases are generated such that digit occurs at least once in number.
*/

func removeDigit(number string, digit byte) string {
	var maxStr string
	for i := range number {
		if number[i] == digit {
			remove := number[:i] + number[i+1:]
			if remove > maxStr {
				maxStr = remove
			}
		}
	}
	return maxStr
}
