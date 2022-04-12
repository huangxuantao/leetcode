package xiaohao

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == 1 {
		if root.Left != nil || root.Right != nil {
			root.Left = pruneTree(root.Left)
			root.Right = pruneTree(root.Right)
			return root
		}

	} else {
		if root.Left != nil || root.Right != nil {
			root.Left = pruneTree(root.Left)
			root.Right = pruneTree(root.Right)

			if root.Left == nil && root.Right == nil {
				root = nil
			}
			return root
		}
		root = nil
	}

	return root
}
