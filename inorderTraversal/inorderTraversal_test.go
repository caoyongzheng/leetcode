package inorderTraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	assert.Equal(t, []int{1}, inorderTraversal(&TreeNode{
		Val: 1,
	}))
	assert.Equal(t, []int{2, 1}, inorderTraversal(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}))
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
}
