package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxDepth(t *testing.T) {
	assert.Equal(t, 3, maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 7},
		},
	}))
}
