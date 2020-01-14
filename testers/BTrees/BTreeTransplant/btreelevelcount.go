package student

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var left = 1 + BTreeLevelCount(root.Left)
	var right = 1 + BTreeLevelCount(root.Right)

	if left > right {
		return left
	}

	return right
}
