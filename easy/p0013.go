package easy

func RomanToInt(s string) int {

	data := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	var pre byte
	for _, c := range []byte(s) {
		if pre > 0 {
			if (pre == 'I' && (c == 'V' || c == 'X')) ||
				(pre == 'X' && (c == 'L' || c == 'C')) ||
				(pre == 'C' && (c == 'D' || c == 'M')) {
				// 表示六种特殊情况
				sum += (data[c] - data[pre])
				pre = 0
			} else {
				sum += data[pre]
				pre = c
			}
		} else {
			pre = c
		}
	}
	if pre > 0 {
		sum += data[pre]
	}
	return sum
}
