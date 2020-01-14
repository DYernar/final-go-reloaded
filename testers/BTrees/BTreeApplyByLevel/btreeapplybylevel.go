package student

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	var level = BTreeLevelCount(root)

	for i := 0; i < level; i++ {
		ApplyByLevelRecursive(root, f, i)
	}
}

func ApplyByLevelRecursive(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}

	if level == 0 {
		f(root.Data)
	}

	ApplyByLevelRecursive(root.Left, f, level-1)
	ApplyByLevelRecursive(root.Right, f, level-1)
}
