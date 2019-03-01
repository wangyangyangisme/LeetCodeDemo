package easy

import "math"

func Reverse(x int) int {
	var ret int64
	for x != 0 {
		ret = ret*10 + int64(x%10)
		x = x / 10
	}

	if ret < math.MinInt32 || ret > math.MaxInt32 {
		return 0
	}
	return int(ret)
}
