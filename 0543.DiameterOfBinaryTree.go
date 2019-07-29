package leetcode_go

func diameterOfBinaryTree(root *TreeNode) int {

	max := 0
	depthAndDiameter(root, &max)

	return max
}

func depthAndDiameter(root *TreeNode, max *int) (depth int) {

	if root == nil {
		return
	}

	leftDepth := depthAndDiameter(root.Left, max)
	rightDepth := depthAndDiameter(root.Right, max)

	if leftDepth > rightDepth {
		depth = leftDepth + 1
	} else {
		depth = rightDepth + 1
	}

	if leftDepth+rightDepth > *max {
		*max = leftDepth + rightDepth
	}

	return
}
