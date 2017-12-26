package leetcode

// ListNode singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	root := &ListNode{}
	for i, pre := head, root; i != nil; {
		pre, root.Next = root.Next, i
		i, root.Next.Next = root.Next.Next, pre
	}
	return root.Next
}
