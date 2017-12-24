package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestHammingDistance(t *testing.T) {
	assert.Equal(t, 2, hammingDistance(1, 4))
}
