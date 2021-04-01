package No47

import (
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{2, 1, 1}
	result := [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}
	res := permuteUnique(nums)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			if res[i][j] != result[i][j] {
				t.Fatalf("Error! 期望值：%v 实际值：%v", result[i][j], res[i][j])
			}
		}
	}
	t.Logf("没毛病！")
}
