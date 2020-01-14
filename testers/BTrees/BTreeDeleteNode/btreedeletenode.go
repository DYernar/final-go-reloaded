package student

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	var parent = node.Parent
	var data = node.Data

	if node.Right != nil {
		BTreeMin(node.Right).Left = node.Left
		node = node.Right
	} else {
		node = node.Left
	}

	if parent != nil {
		if parent.Data <= data {
			parent.Right = node
		} else {
			parent.Left = node
		}
	} else {
		root = node
	}
	node.Parent = parent

	return root
}
