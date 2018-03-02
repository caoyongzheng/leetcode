package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubSets(t *testing.T) {
	assert.Equal(t, [][]int{}, subsets([]int{1, 2, 3}))
}
