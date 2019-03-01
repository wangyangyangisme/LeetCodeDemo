package easy

func RemoveDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			copy(nums[i:], nums[i+1:])
			i--
			l--
		}
	}
	return l
}
