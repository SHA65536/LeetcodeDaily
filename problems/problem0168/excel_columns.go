package problem0168

/*
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
*/

func convertToTitle(columnNumber int) string {
	var res []byte
	for columnNumber > 0 {
		columnNumber--
		t := columnNumber % 26
		res = append(res, 'A'+byte(t))
		columnNumber /= 26
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
