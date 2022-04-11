package xiaohao

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxFunc(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxFunc(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
