package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestDiameterOfBinaryTree(t *testing.T) {
	assert.Equal(t, 3, diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}
