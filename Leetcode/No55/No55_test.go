package No55

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	flag := canJump(nums)
	fmt.Println("bool: ", flag)
}
