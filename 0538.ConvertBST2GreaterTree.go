package leetcode_go

func convertBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var sum int

	convert(root, &sum)

	return root
}

func convert(root *TreeNode, sum *int) {

	if root.Right != nil {
		convert(root.Right, sum)
	}

	root.Val += *sum
	*sum = root.Val

	if root.Left != nil {
		convert(root.Left, sum)
	}
}
