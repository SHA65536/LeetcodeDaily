package problem0989

/*
The array-form of an integer num is an array representing its digits in left to right order.
    For example, for num = 1321, the array form is [1,3,2,1].
Given num, the array-form of an integer, and an integer k, return the array-form of the integer num + k.
*/

func addToArrayForm(num []int, k int) []int {
	var carry, temp int
	// Reverse the array
	reverse(num)
	// Add until the array is exhausted
	for i := range num {
		temp = (num[i] + (k % 10) + carry) % 10
		carry = (num[i] + (k % 10) + carry) / 10
		num[i] = temp
		k /= 10
	}
	// Add until k and carry are exhausted
	for carry > 0 || k > 0 {
		// We are past the original array so we got to append
		num = append(num, ((k%10)+carry)%10)
		carry = ((k % 10) + carry) / 10
		k /= 10
	}
	// Reverse again for numerical order
	reverse(num)
	return num
}

func reverse(num []int) {
	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-i-1] = num[len(num)-i-1], num[i]
	}
}
