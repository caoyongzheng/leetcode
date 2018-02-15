package countBits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2}, countBits(5))
}
