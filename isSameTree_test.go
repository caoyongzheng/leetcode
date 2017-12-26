package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestIsSameTree(t *testing.T) {
	assert.Equal(t, true, isSameTree(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}, &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))

	assert.Equal(t, false, isSameTree(&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}, &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}))

	assert.Equal(t, false, isSameTree(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}, &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}))

	assert.Equal(t, false, isSameTree(&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: -5},
	}, &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: -8},
	}))
}
