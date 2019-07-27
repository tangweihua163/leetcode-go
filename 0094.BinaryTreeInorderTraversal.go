package leetcode_go

func inorderTraversal(root *TreeNode) []int {

	result := make([]int, 0, 1000)

	if root == nil {
		return result
	}

	result = inorderTraversal2Slice(root, result)

	return result
}

func inorderTraversal2Slice(root *TreeNode, s []int) []int {

	if root.Left != nil {
		s = inorderTraversal2Slice(root.Left, s)
	}

	s = append(s, root.Val)

	if root.Right != nil {
		s = inorderTraversal2Slice(root.Right, s)
	}

	return s
}
