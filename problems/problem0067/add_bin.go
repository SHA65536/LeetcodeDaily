package problem0067

/*
Given two binary strings a and b, return their sum as a binary string.
*/

func addBinary(a string, b string) string {
	var res []byte = make([]byte, 0, max(len(a), len(b))+1)
	var cura, curb int
	var carry bool
	var temp byte
	cura, curb = len(a)-1, len(b)-1
	for cura >= 0 || curb >= 0 {
		if cura < 0 {
			temp, carry = add('0', b[curb], carry)
		} else if curb < 0 {
			temp, carry = add(a[cura], '0', carry)
		} else {
			temp, carry = add(a[cura], b[curb], carry)
		}
		res = append(res, temp)
		cura--
		curb--
	}
	if carry {
		res = append(res, '1')
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}

func add(a, b byte, carry bool) (byte, bool) {
	if a == '1' && b == '1' {
		if carry {
			return '1', true
		}
		return '0', true
	} else if a == '1' || b == '1' {
		if carry {
			return '0', true
		}
		return '1', false
	} else {
		if carry {
			return '1', false
		}
		return '0', false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
