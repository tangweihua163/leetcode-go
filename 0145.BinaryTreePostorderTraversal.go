package leetcode_go

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	result := make([]int, 0, 1000)
	stack := make([]*TreeNode, 0, 1000)

	var pre, cur *TreeNode

	stack = append(stack, root)

	for len(stack) > 0 {

		cur = stack[len(stack)-1]

		if cur.Left == nil && cur.Right == nil ||
			(pre != nil) && (pre == cur.Left || pre == cur.Right) {

			result = append(result, cur.Val)
			stack = stack[:len(stack)-1]
			pre = cur

			continue
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}

	return result
}
