package easy

func MaxProfit2(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	buy := 0
	profit := 0
	isHave := false
	for i := 0; i < l-1; i++ {
		curPrice := prices[i]
		nextPrice := prices[i+1]
		if nextPrice > curPrice {
			// 明天价格比今天高
			if isHave {
				// 当前有股票，继续持有
			} else {
				// 当前没有股票
				isHave = true
				buy = curPrice
			}
		} else if nextPrice < curPrice {
			// 明天价格低于今天价格
			if isHave {
				// 当前有股票，卖出
				profit += (curPrice - buy)
				isHave = false
			} else {
				// 当前没有股票
			}
		}
	}
	if isHave {
		profit += (prices[l-1] - buy)
	}
	return profit
}
