package leetcode_go

func pathSum(root *TreeNode, sum int) int {
	stack := make([]int, 1000)
	var count int

	pathSumImpl(root, sum, stack, &count)

	return count
}

func pathSumImpl(root *TreeNode, sum int, stack []int, count *int) {
	if root == nil {
		return
	}

	stack = append(stack, root.Val)
	var s int
	for i := len(stack) - 1; i >= 0; i-- {
		s += stack[i]
		if s == sum {
			*count++
		}
	}

	pathSumImpl(root.Left, sum, stack, count)
	pathSumImpl(root.Right, sum, stack, count)

	stack = stack[:len(stack)-1]
}
