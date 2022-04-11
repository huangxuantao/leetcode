package xiaohao

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if level == len(*res) {
		*res = append(*res, []int{root.Val})
	} else {
		(*res)[level] = append((*res)[level], root.Val)
	}

	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)

	return
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	dfs(root, 0, &res)
	return res
}