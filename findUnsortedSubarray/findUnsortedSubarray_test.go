package findUnsortedSubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUnsortedSubarray(t *testing.T) {
	assert.Equal(t, 0, findUnsortedSubarray(nil))
	assert.Equal(t, 0, findUnsortedSubarray([]int{}))
	assert.Equal(t, 0, findUnsortedSubarray([]int{1}))
	assert.Equal(t, 2, findUnsortedSubarray([]int{2, 1}))
	assert.Equal(t, 5, findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
