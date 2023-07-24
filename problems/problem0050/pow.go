package problem0050

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/

func myPow(x float64, n int) float64 {
	var pow float64 = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n != 0 {
		if n&1 != 0 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
