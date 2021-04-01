package No46

func permute(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return make([][]int, 0)
	}
	var result [][]int = make([][]int, 0)
	var flag []bool = make([]bool, len(nums))
	var res []int = make([]int, 0)
	result = dfs(nums, length, 0, flag, res, result)
	return result
}

func dfs(nums []int, len int, depth int, flag []bool, res []int, result [][]int) [][]int {
	if depth == len {
		temp := make([]int, len)
		copy(temp, res)
		result = append(result, temp)
		return result
	}
	for i := 0; i < len; i++ {
		if flag[i] == false {
			res = append(res, nums[i])
			flag[i] = true
			result = dfs(nums, len, depth+1, flag, res, result)
			flag[i] = false
			res = res[:depth]
		}
	}
	return result
}
