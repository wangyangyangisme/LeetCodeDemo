package easy

func MaxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min := prices[0]
	max := 0
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > max {
			max = v - min
		}
	}
	return max
}
