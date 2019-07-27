package leetcode_go

func isSymmetric(root *TreeNode) bool {

	level := make([]*TreeNode, 0, 1)
	level = append(level, root)
	var start, end, startI, endI int

	for {
		startI = start
		endI = end

		for startI < endI {
			if level[startI] == nil && level[endI] == nil ||
				level[startI] != nil && level[endI] != nil &&
					level[startI].Val == level[endI].Val {
				startI++
				endI--
				continue
			}
			return false
		}

		nextLevel := make([]*TreeNode, 0, len(level)*2)
		for i := start; i <= end; i++ {
			if level[i] != nil {
				nextLevel = append(nextLevel, level[i].Left)
				nextLevel = append(nextLevel, level[i].Right)
			}
		}

		if len(nextLevel) == 0 {
			break
		}

		level = nextLevel

		start = 0
		end = len(nextLevel) - 1
	}

	return true
}
