package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 21, maxSubArray([]int{8, -19, 5, -4, 20}))
	assert.Equal(t, 4, maxSubArray([]int{3, -2, -3, -3, 1, 3, 0}))
	assert.Equal(t, 3, maxSubArray([]int{1, -3, 2, 0, -1, 0, -2, -3, 1, 2, -3, 2}))
	assert.Equal(t, 46, maxSubArray([]int{8, -1, -9, 3, 8, 1, -4, 8, 6, -1, 8, 4, 2, 4, 7, -5}))
}
