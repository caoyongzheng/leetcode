package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, 2, minDistance("abcd", "bcda"))
	assert.Equal(t, 1, minDistance("abcda", "bcda"))
	assert.Equal(t, 1, minDistance("a", "b"))
	// assert.Equal(t, 7, minDistance("dinitrophenylhydrazine", "benzalphenylhydrazone"))
	assert.Equal(t, 3, minDistance("park", "spake"))
	assert.Equal(t, 3, minDistance("mart", "karma"))
	// assert.Equal(t, 3, minDistance("trinitrophenylmethylnitraminemart", "dinitrophenylhydrazine"))
}
