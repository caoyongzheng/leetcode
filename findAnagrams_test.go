package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	assert.Equal(t, []int{0, 6}, findAnagrams("cbaebabacd", "abc"))
}
