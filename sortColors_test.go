package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestSortColors(t *testing.T) {
	c1 := []int{2, 2, 1, 1, 0, 0}
	sortColors(c1)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, c1)

	c2 := []int{1, 2, 1, 1, 0, 0}
	sortColors(c2)
	assert.Equal(t, []int{0, 0, 1, 1, 1, 2}, c2)
}
