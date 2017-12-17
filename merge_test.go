package leetcode

import "testing"
import "github.com/stretchr/testify/assert"

func TestMerge(t *testing.T) {
	assert.Equal(t, []Interval{
		Interval{Start: 1, End: 6},
		Interval{Start: 8, End: 10},
		Interval{Start: 15, End: 18},
	}, merge([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 2, End: 6},
		Interval{Start: 8, End: 10},
		Interval{Start: 15, End: 18},
	}))

	assert.Equal(t, []Interval{
		Interval{Start: 1, End: 4},
	}, merge([]Interval{
		Interval{Start: 1, End: 4},
		Interval{Start: 2, End: 3},
	}))
}
