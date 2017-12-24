package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{1, 2, 2}))
}
