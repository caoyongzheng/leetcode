package minStack

// 155. Min Stack
// description: https://leetcode.com/problems/min-stack/description/

// MinStack 最小栈堆
type MinStack struct {
	elements []int
}

// Constructor MinStack construct func
func Constructor() MinStack {
	return MinStack{}
}

// Push ...
func (m *MinStack) Push(x int) {
	m.elements = append(m.elements, x)
}

// Pop ...
func (m *MinStack) Pop() {
	if len(m.elements) == 0 {
		return
	}
	m.elements = m.elements[:len(m.elements)-1]
}

// Top ...
func (m *MinStack) Top() int {
	if len(m.elements) == 0 {
		return 0
	}
	return m.elements[len(m.elements)-1]
}

// GetMin ...
func (m *MinStack) GetMin() int {
	if len(m.elements) == 0 {
		return 0
	}
	minElement := m.elements[0]
	for _, v := range m.elements {
		if minElement > v {
			minElement = v
		}
	}
	return minElement
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
