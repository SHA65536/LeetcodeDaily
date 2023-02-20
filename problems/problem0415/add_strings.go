package problem0415

/*
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
You must also not convert the inputs to integers directly.
*/

func addStrings(num1 string, num2 string) string {
	var carry, idx, temp int
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	var A, B = []byte(num1), []byte(num2)
	reverse(A)
	reverse(B)
	for idx = 0; idx < len(A); idx++ {
		temp = int(A[idx]-'0') + int(B[idx]-'0') + carry
		carry = temp / 10
		A[idx] = byte(temp%10) + '0'
	}
	for ; idx < len(B); idx++ {
		temp = int(B[idx]-'0') + carry
		carry = temp / 10
		A = append(A, byte(temp%10)+'0')
	}
	if carry > 0 {
		A = append(A, byte(carry)+'0')
	}
	reverse(A)
	return string(A)
}

func reverse(num []byte) {
	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-i-1] = num[len(num)-i-1], num[i]
	}
}
