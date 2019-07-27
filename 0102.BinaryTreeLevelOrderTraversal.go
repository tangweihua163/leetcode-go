package leetcode_go

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0, 100)

	if root == nil {
		return result
	}

	level := make([]*TreeNode, 0, 1)
	level = append(level, root)

	for {

		row := make([]int, 0, len(level))

		for _, p := range level {
			row = append(row, p.Val)
		}

		result = append(result, row)

		nextLevel := make([]*TreeNode, 0, len(level)*2)
		for _, p := range level {
			if p.Left != nil {
				nextLevel = append(nextLevel, p.Left)
			}
			if p.Right != nil {
				nextLevel = append(nextLevel, p.Right)
			}
		}

		if len(nextLevel) == 0 {
			break
		}

		level = nextLevel
	}

	return result
}
