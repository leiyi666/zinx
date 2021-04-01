package No123

import (
	"sort"
)

/**
不会
*/

func maxProfit(prices []int) int {
	var profit = make([]int, 0)
	cur := prices[0]
	n := len(prices)
	temp := 0
	for i := 0; i < n; i++ {
		cur = prices[i]
		temp = nextMax(prices, i, n)
		if prices[temp]-cur > 0 {
			profit = append(profit, prices[temp]-cur)
		}
		i = temp
	}
	switch len(profit) {
	case 0:
		return 0
	case 1:
		return profit[0]
	case 2:
		return profit[0] + profit[1]
	default:
		{
			sort.Sort(sort.Reverse(sort.IntSlice(profit)))
			return profit[0] + profit[1]
		}
	}
}

func nextMax(prices []int, index int, length int) int {
	for i := index + 1; i < length; i++ {
		if prices[i] < prices[i-1] {
			return i - 1
		}
	}
	return length - 1
}
