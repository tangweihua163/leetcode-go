package leetcode_go

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Next   *Node
	Random *Node
}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	level := make([]*Node, 0, 1)
	level = append(level, root)

	for len(level) != 0 {
		for i := len(level) - 1; i >= 0; i-- {
			if i == len(level)-1 {
				level[i].Next = nil
			} else {
				level[i].Next = level[i+1]
			}
		}

		newLevel := make([]*Node, 0, len(level)*2)
		for i := 0; i < len(level); i++ {
			if level[i].Left == nil {
				break
			}
			newLevel = append(newLevel, level[i].Left, level[i].Right)
		}

		level = newLevel
	}

	return root
}
