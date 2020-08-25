package leetcode_go

func inorderTraversal(root *TreeNode) []int {

	result := make([]int, 0, 1000)
	stack := make([]*TreeNode, 0, 1000)

	var p = root

	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			result = append(result, p.Val)
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}

	return result
}
