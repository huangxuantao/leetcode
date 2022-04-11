package xiaohao

import "math"

func isValidBST(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}
