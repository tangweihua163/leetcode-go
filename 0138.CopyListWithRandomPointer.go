package leetcode_go

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	ma := make(map[*Node]*Node)
	mb := make(map[*Node]*Node)

	for p := head; p != nil; p = p.Next {
		ma[p] = p.Random
	}

	var l, t *Node
	for p := head; p != nil; p = p.Next {
		node := &Node{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}

		if l == nil {
			l = node
			t = node
		} else {
			t.Next = node
			t = t.Next
		}

		mb[p] = node
	}

	for k, v := range ma {
		mb[k].Random = mb[v]
	}

	return l
}
