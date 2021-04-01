package No55

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	maxNum := 0
	for i := 0; i < n; i++ {
		if maxNum >= n-1 {
			return true
		}
		if i > maxNum {
			return false
		}
		maxNum = max(maxNum, i+nums[i])
	}
	return true
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
