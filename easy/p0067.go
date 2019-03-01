package easy

func AddBinary(a string, b string) string {
	ret := []byte{}
	la := len(a)
	lb := len(b)
	ax, bx, cx, x := byte('0'), byte('0'), byte('0'), byte('0')
	for i := 1; i <= la || i <= lb; i++ {
		if (la - i) < 0 {
			ax = '0'
		} else {
			ax = a[la-i]
		}
		if (lb - i) < 0 {
			bx = '0'
		} else {
			bx = b[lb-i]
		}
		cx, x = add(ax, bx, cx)
		ret = append([]byte{x}, ret[0:]...)
	}
	if cx-'0' > 0 {
		ret = append([]byte{cx}, ret[0:]...)
	}
	return string(ret)
}

/*
 a,b为相加的每个位上的数值
 c为进位数值
*/
func add(a byte, b byte, c byte) (byte, byte) {
	sum := (a - '0') + (b - '0') + (c - '0')
	switch sum {
	case 0:
		return '0', '0'
	case 1:
		return '0', '1'
	case 2:
		return '1', '0'
	case 3:
		return '1', '1'
	default:
		return '0', '0'
	}
	return '0', '0'
}
