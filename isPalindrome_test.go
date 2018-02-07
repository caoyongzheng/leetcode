package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome(&ListNode{
		Val: 1,
	}))
	assert.Equal(t, false, isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}))
	assert.Equal(t, true, isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}))
}
