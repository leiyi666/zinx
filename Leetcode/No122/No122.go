package No122

func maxProfit(prices []int) int {
	n := len(prices)
	cur := prices[0]
	price := 0
	for i := 0; i < n; i++ {
		if prices[i] > cur {
			price = price + (prices[i] - cur)
		}
		cur = prices[i]
	}
	return price
}
