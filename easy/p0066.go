package easy

func PlusOne(digits []int) []int {
	isPlus := true
	for i := len(digits) - 1; i >= 0; i-- {
		if isPlus {
			digits[i]++
		}
		if digits[i] == 10 {
			digits[i] = 0
			isPlus = true
		} else {
			isPlus = false
		}
	}
	if isPlus {
		//digits = append(digits[:0], append([]int{1}, digits[0:]...)...)
		digits = append([]int{1}, digits[:]...)
	}
	return digits
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits[:]...)
	} else {
		return digits
	}
}
