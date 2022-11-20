package problem0224

/*
Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.
Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().
*/

func calculate(s string) int {
	var stack = []int{} // We will use this slice a stack for order of operations
	var curNum, result, sign int
	sign = 1 // Sign will be 1 or -1 depending on positive or negative

	for i := range s {
		// If we see a digit
		if s[i] >= '0' && s[i] <= '9' {
			// Add it to the end of curNum
			curNum = 10*curNum + int(s[i]-'0')
			continue
		}

		switch s[i] {
		case '+':
			result += sign * curNum // Execute last operation
			sign, curNum = 1, 0     // Create a new positive number
		case '-':
			result += sign * curNum // Execute last operation
			sign, curNum = -1, 0    // Create a new negative number
		case '(':
			stack = append(stack, result, sign) // Stash the result and sign so far
			sign, result = 1, 0
		case ')':
			result += sign * curNum // Execute last operation
			curNum = 0
			result *= stack[len(stack)-1] // Apply stashed sign
			result += stack[len(stack)-2] // Apply stashed result
			stack = stack[:len(stack)-2]  // Remove from stack
		}
	}
	if curNum != 0 { // Execute last operation if there is one
		result += sign * curNum
	}
	return result
}
