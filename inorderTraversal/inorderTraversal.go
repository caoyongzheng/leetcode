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
	for v := root; l != nil && v != nil; {
		if v.Left != nil {
			l = &Link{Node: v.Left, Next: l}
			v = l.Node
		} else {
			r = append(r, v.Val)
			v = v.Right
			for l = l.Next; v == nil; l = l.Next {
				if l == nil {
					return
				}
				v = l.Node
				r = append(r, v.Val)
				v = v.Right
			}
			l = &Link{Node: v, Next: l}
		}
	}
	return
}
