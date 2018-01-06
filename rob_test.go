package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestRob(t *testing.T) {
	assert.Equal(t, 0, rob([]int{}))
	assert.Equal(t, 2, rob([]int{2}))
	assert.Equal(t, 4, rob([]int{2, 4}))
	assert.Equal(t, 7, rob([]int{2, 4, 5}))
	assert.Equal(t, 8, rob([]int{2, 4, 5, 4}))
	assert.Equal(t, 10, rob([]int{2, 4, 5, 4, 3}))
}
