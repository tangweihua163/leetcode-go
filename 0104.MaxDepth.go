package leetcode_go

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth < rightDepth {
		return rightDepth + 1
	} else {
		return leftDepth + 1
	}
}
