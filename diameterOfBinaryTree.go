package leetcode

func max3(a, b, c int) int {
	return max(a, max(b, c))
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := heightOfBinaryTree(root)
	return diameter
}

func heightOfBinaryTree(root *TreeNode) (height int, diameter int) {
	if root == nil {
		return 0, 0
	}
	leftHeight, leftDiameter := heightOfBinaryTree(root.Left)
	rightHeight, rightDiameter := heightOfBinaryTree(root.Right)
	height, diameter = max(leftHeight, rightHeight)+1, max3(leftDiameter, rightDiameter, leftHeight+rightHeight)
	return
}
