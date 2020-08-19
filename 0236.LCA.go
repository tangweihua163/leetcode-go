package leetcode_go

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p == q {
		return p
	}

	ss := make([][]*TreeNode, 2)
	s := make([]*TreeNode, 10000)

	ss, s = findLCA(root, p, q, ss, s)

	if len(ss) != 2 {
		return nil
	}

	var i = 0
	for i = 0; i < len(ss[0]) && i < len(ss[1]); i++ {
		if ss[0][i] != ss[1][i] {
			break
		}
	}

	return ss[0][i-1]

}

func findLCA(root, p, q *TreeNode, ss [][]*TreeNode, s []*TreeNode) ([][]*TreeNode, []*TreeNode) {

	if len(ss) == 2 {
		return ss, s
	}

	if root == nil {
		return ss, s
	}

	s = append(s, root)
	if root == p || root == q {
		dst := make([]*TreeNode, len(s))
		copy(dst, s)
		ss = append(ss, dst)
	}

	if len(ss) == 2 {
		return ss, s
	}

	ss, s = findLCA(root.Left, p, q, ss, s)
	if len(ss) == 2 {
		return ss, s
	}

	ss, s = findLCA(root.Right, p, q, ss, s)
	if len(ss) == 2 {
		return ss, s
	}

	s = s[:len(s)-1]

	return ss, s
}
