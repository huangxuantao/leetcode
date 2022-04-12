package xiaohao

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Right == nil {
		return root.Left
	}

	if root.Left == nil {
		return root.Right
	}

	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	root.Val = minNode.Val
	root.Right = deleteMinNode(root.Right)

	return root
}

func deleteMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root.Right
	}

	root.Left = deleteMinNode(root.Left)
	return root
}
