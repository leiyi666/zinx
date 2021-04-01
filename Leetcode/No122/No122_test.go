package No122

import "testing"

func TestMaxProfit(t *testing.T) {
	//prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1, 2, 3, 4, 5}
	prices := []int{7, 6, 4, 3, 1}
	result := maxProfit(prices)
	if result != 0 {
		t.Fatalf("不对劲！期望值：7 实际值：%v", result)
	}
	t.Logf("没毛病!")
}
