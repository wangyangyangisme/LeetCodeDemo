package easy

func RemoveElement(nums []int, val int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		if nums[i] == val {
			continue
		}
		if res != i {
			nums[res] = nums[i]
		}
		res++
	}
	return res
}
