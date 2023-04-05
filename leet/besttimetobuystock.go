package leet

func MaxProfit(prices []int) int {
	b := prices[0]
	var diff int
	for i := 1; i < len(prices); i++ {
		if prices[i]-b < 0 {
			b = prices[i]
		} else if prices[i]-b > 0 {
			diff = max(diff, prices[i]-b)
		}
	}

	return diff
}
