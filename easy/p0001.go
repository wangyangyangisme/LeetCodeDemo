package easy

func TwoSum(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int, l)
	for i := 0; i < l; i++ {
		if j, ok := m[target-nums[i]]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}
