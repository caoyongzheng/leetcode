package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestConvertBST(t *testing.T) {
	assert.Equal(t, &TreeNode{
		Val: 18,
		Left: &TreeNode{
			Val: 20,
		},
		Right: &TreeNode{
			Val: 13,
		},
	}, convertBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 13,
		},
	}))
}
