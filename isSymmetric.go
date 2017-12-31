package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	invertTree(root.Left)
	return isSameTree(root.Left, root.Right)
}
