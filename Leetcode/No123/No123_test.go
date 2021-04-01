package No123

import "testing"

func TestMaxProfit(t *testing.T) {
	//prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	prices := []int{1, 2, 3, 4, 5}
	//prices := []int{7, 6, 4, 3, 1}
	//prices := []int{6, 1, 3, 2, 4, 7}
	//prices := []int{3, 2, 6, 5, 0, 3}
	result := maxProfit(prices)
	if result != 6 {
		t.Fatalf("不对劲！ 期望值：6 实际值：%v", result)
	}
	t.Logf("没毛病！")
}
