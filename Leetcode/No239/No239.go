package No239

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}
	queue := make([]int, 0)
	result := make([]int, 0)
	length := 0
	for i := 0; i < len(nums); i++ {
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
			length--
		}
		for len(queue) != 0 && nums[i] > queue[length-1] {
			queue = queue[:length-1]
			length--
		}
		queue = append(queue, nums[i])
		length++
		if i >= k-1 {
			result = append(result, queue[0])
		}
	}
	return result
}
