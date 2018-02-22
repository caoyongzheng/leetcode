package inorderTraversal

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Link ...
type Link struct {
	Node *TreeNode
	Next *Link
}

func inorderTraversal(root *TreeNode) (r []int) {
	l := &Link{Node: root}
	for l.Node != nil {
		node := l.Node
		if node.Left != nil {
			l = &Link{Node: node.Left, Next: l}
			continue
		}

		r = append(r, node.Val)
		for l = l.Next; node.Right == nil; l = l.Next {
			if l == nil {
				return
			}
			node = l.Node
			r = append(r, node.Val)
		}
		l = &Link{Node: node.Right, Next: l}
	}
	return
}
