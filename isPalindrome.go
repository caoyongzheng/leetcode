package leetcode

// 234. Palindrome Linked List
// Given a singly linked list, determine if it is a palindrome.

// Follow up:
// Could you do it in O(n) time and O(1) space?

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}

	for slow := reverseList(slow); slow != nil; slow, head = slow.Next, head.Next {
		if slow.Val != head.Val {
			return false
		}
	}

	return true
}
