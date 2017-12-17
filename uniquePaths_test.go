package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 1, uniquePaths(1, 2))
	assert.Equal(t, 2, uniquePaths(2, 2))
	assert.Equal(t, 1916797311, uniquePaths(9, 51))
}
