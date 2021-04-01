package No98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(temp *TreeNode, lower int, upper int) bool {
	if temp == nil {
		return true
	}
	if temp.Val <= lower || temp.Val >= upper {
		return false
	}
	return helper(temp.Left, lower, temp.Val) && helper(temp.Right, temp.Val, upper)
}
