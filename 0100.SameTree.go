package leetcode_go

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == q {
		return true
	}

	if p != nil && q == nil {
		return false
	}
	if q != nil && p == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
