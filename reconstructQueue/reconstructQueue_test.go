package reconstructQueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReconstructQueue(t *testing.T) {
	assert.Equal(t, [][]int{}, reconstructQueue(nil))
	assert.Equal(t, [][]int{
		[]int{5, 0}, []int{7, 0}, []int{5, 2}, []int{6, 1}, []int{4, 4}, []int{7, 1},
	}, reconstructQueue([][]int{
		[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2},
	}))
}
