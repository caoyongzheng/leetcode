package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestMoveZeroes(t *testing.T) {
	data1 := []int{0, 1, 0, 3, 12}
	moveZeroes(data1)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, data1)

	data2 := []int{2, 1}
	moveZeroes(data2)
	assert.Equal(t, []int{2, 1}, data2)
}
