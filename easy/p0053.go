package easy

// O(n)
func MaxSubArray2(nums []int) int {
	l := len(nums)
	if l < 1 {
		return 0
	}
	max, temp := nums[0], 0
	for _, v := range nums {
		if temp < 0 {
			//如果当前子列和为负数,则不可能使后面的部分和增大，抛弃之
			temp = v
		} else {
			temp += v
		}
		if temp > max {
			//发现更大和，则更新当前结果
			max = temp
		}
	}
	return max
}

// O(n^2)
func MaxSubArray(nums []int) int {
	l := len(nums)
	if l < 1 {
		return 0
	}
	max, temp := nums[0], 0
	for i := range nums {
		temp = 0
		for _, v2 := range nums[i:] {
			temp += v2
			if temp > max {
				max = temp
			}
		}
	}
	return max
}
