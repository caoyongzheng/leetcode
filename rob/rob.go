package rob

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	m := map[*TreeNode]int{nil: 0}
	return rob2(root, m)
}

func rob2(root *TreeNode, m map[*TreeNode]int) int {
	if v, ok := m[root]; ok {
		return v
	}

	withRoot := root.Val
	if root.Left != nil {
		withRoot += rob2(root.Left.Left, m) + rob2(root.Left.Right, m)
	}
	if root.Right != nil {
		withRoot += rob2(root.Right.Left, m) + rob2(root.Right.Right, m)
	}

	count := max(withRoot, rob2(root.Left, m)+rob2(root.Right, m))
	m[root] = count
	return count
}
