package easy

func MySqrt(x int) int {
	for i := 0; i <= x; i++ {
		tmp := i * i
		if tmp == x {
			return i
		} else if tmp > x {
			return i - 1
		}
	}
	return 0
}

func MySqrt2(x int) int {
	var i, j, mid, tmp int = 0, x, 0, 0

	for i = 0; i < j; {
		mid = i + (j-i)/2
		tmp = mid * mid

		if tmp == x {
			return mid
		} else if tmp > x {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	if i*i > x {
		return i - 1
	}
	return i
}
