package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var t = new(TreeNode)
	t.Val = nums[len(nums)/2]
	t.Left = sortedArrayToBST(nums[:len(nums)/2])
	t.Right = sortedArrayToBST(nums[len(nums)/2+1:])

	return t
}
