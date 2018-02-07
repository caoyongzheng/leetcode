package minStack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	m := Constructor()
	assert.Equal(t, 0, m.GetMin())
	m.Push(-2)
	m.Push(1)
	m.Push(-3)
	assert.Equal(t, -3, m.GetMin())
	m.Pop()
	assert.Equal(t, 1, m.Top())
	assert.Equal(t, -2, m.GetMin())
}
