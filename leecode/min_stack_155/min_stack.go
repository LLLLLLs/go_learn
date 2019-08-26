// Time        : 2019/07/31
// Description :

package min_stack_155

import "math"

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// push(x) -- Push element x onto stack.
// pop() -- Removes the element on top of the stack.
// top() -- Get the top element.
// getMin() -- Retrieve the minimum element in the stack.
//
//
// Example:
//
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> Returns -3.
// minStack.pop();
// minStack.top();      --> Returns 0.
// minStack.getMin();   --> Returns -2.

type MinStack struct {
	values []int
	min    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		values: make([]int, 0),
		min:    math.MaxInt32,
	}
}

func (s *MinStack) Push(x int) {
	s.values = append(s.values, x)
	if x < s.min {
		s.min = x
	}
}

func (s *MinStack) Pop() {
	p := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	if p == s.min {
		s.refreshMin()
	}
}

func (s *MinStack) refreshMin() {
	s.min = math.MaxInt32
	for i := range s.values {
		if s.values[i] < s.min {
			s.min = s.values[i]
		}
	}
}

func (s *MinStack) Top() int {
	return s.values[len(s.values)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}
