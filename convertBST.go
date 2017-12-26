package leetcode

// Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

// Example:

// Input: The root of a Binary Search Tree like this:
//               5
//             /   \
//            2     13

// Output: The root of a Greater Tree like this:
//              18
//             /   \
//           20     13

func walkBST(root *TreeNode, fn func(node *TreeNode)) {
	if root != nil {
		walkBST(root.Right, fn)
		fn(root)
		walkBST(root.Left, fn)
	}
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	walkBST(root, func(node *TreeNode) {
		sum += node.Val
		node.Val = sum
	})
	return root
}
