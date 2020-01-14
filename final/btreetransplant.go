package student

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	parent := node.Parent
	if parent == nil {
		return rplc
	}
	if parent.Left == node {
		parent.Left = rplc
		rplc.Parent = parent
		return root
	}

	if parent.Right == node {
		parent.Right = rplc
		rplc.Parent = parent
		return root
	}
	return root
}
