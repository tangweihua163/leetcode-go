package leetcode_go

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 10000)
	stack := make([]*TreeNode, 0, 10000)

	var p = root

	for p != nil || len(stack) > 0 {

		for p != nil {
			result = append(result, p.Val)
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}

	}

	return result
}
