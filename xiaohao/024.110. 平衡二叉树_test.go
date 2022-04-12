package xiaohao

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	leftH := maxDepth(root.Left) + 1
	rightH := maxDepth(root.Right) + 1
	if math.Abs(float64(rightH)-float64(leftH)) > 1 {
		return false
	}

	return true
}
