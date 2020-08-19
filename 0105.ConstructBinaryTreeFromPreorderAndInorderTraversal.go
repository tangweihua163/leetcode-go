package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	node := new(TreeNode)
	node.Right = nil
	node.Left = nil
	node.Val = preorder[0]

	if len(preorder) == 1 {
		return node
	}

	var i = 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	if i == len(inorder) {
		return nil
	}

	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return node
}
