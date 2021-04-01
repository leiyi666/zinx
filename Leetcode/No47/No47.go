package No47

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	length := len(nums)
	var res []int = make([]int, 0)
	var result [][]int = make([][]int, 0)
	var flag []bool = make([]bool, length)
	sort.Ints(nums)
	result = dfs(nums, 0, length, flag, res, result)
	return result
}

func dfs(nums []int, depth int, length int, flag []bool, res []int, result [][]int) [][]int {
	if depth == length {
		var temp []int = make([]int, length)
		copy(temp, res)
		result = append(result, temp)
		return result
	}
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] && flag[i-1] == false {
			continue
		}
		if flag[i] == false {
			res = append(res, nums[i])
			flag[i] = true
			result = dfs(nums, depth+1, length, flag, res, result)
			flag[i] = false
			res = res[:depth]
		}
	}
	return result
}
