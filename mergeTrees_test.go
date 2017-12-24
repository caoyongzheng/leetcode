package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestMergeTrees(t *testing.T) {
	assert.Equal(t, &TreeNode{
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
	}, mergeTrees(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}, &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}
