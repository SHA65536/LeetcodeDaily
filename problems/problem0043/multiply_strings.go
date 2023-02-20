package problem0043

/*
Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.
Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

func multiply(num1 string, num2 string) string {
	var res []byte
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// If we operate on the smaller one, we'll do less operations
	A, B := []byte(num1), []byte(num2)
	if len(A) > len(B) {
		A, B = B, A
	}
	// Reverse order for easy manipulation
	reverse(A)
	reverse(B)

	// Doing long multiplication
	for i := range A {
		// Putting extra 0's at the start of the number (end after reverse)
		var row = makeEmpty(i)
		var carry, temp int
		// Multiplying every digits
		for j := range B {
			temp = int(A[i]-'0')*int(B[j]-'0') + carry
			carry = temp / 10
			row = append(row, byte(temp%10)+'0')
		}
		// Remember to carry
		if carry > 0 {
			row = append(row, byte(carry)+'0')
		}
		// Adding to res
		res = add(res, row)
	}

	reverse(res)
	return string(res)
}

// add adds two numbers represented as reversed bytes
func add(num1, num2 []byte) []byte {
	var carry, idx, temp int
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	for idx = 0; idx < len(num1); idx++ {
		temp = int(num1[idx]-'0') + int(num2[idx]-'0') + carry
		carry = temp / 10
		num1[idx] = byte(temp%10) + '0'
	}
	for ; idx < len(num2); idx++ {
		temp = int(num2[idx]-'0') + carry
		carry = temp / 10
		num1 = append(num1, byte(temp%10)+'0')
	}
	if carry > 0 {
		num1 = append(num1, byte(carry)+'0')
	}
	return num1
}

func reverse(num []byte) {
	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-i-1] = num[len(num)-i-1], num[i]
	}
}

func makeEmpty(l int) []byte {
	var res = make([]byte, l)
	for i := range res {
		res[i] = '0'
	}
	return res
}
