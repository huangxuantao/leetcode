package xiaohao

import "math"

func countNodes(root *TreeNode) int {
	d := maxDepth(root)

	var v int
	countDepthNodes(root, 0, d-1, &v)

	for i := 0; i < d-1; i++ {
		v += int(math.Pow(2, float64(i)))
	}
	return v
}

func countDepthNodes(root *TreeNode, level int, depth int, count *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil && depth == level {
		*count++
		return
	}

	countDepthNodes(root.Left, level+1, depth, count)
	countDepthNodes(root.Right, level+1, depth, count)
}

func countNodes2(root *TreeNode) int {
	if root != nil {
		return 1 + countNodes2(root.Left) + countNodes2(root.Right)
	}
	return 0
}
