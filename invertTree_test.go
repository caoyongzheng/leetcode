package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestInvertTree(t *testing.T) {
	assert.Equal(t, &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
	}, invertTree(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}))
}
