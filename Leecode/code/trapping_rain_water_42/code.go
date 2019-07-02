// Time        : 2019/07/02
// Description :

package trapping_rain_water_42

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//
//
// The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
// In this case, 6 units of rain water (blue section) are being trapped.
//
// Example:
//
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6

type stack struct {
	list  []int
	index int
}

func newStack() *stack {
	return &stack{
		list:  make([]int, 0),
		index: 0,
	}
}

func (s *stack) pop() {
	s.list = s.list[:len(s.list)-1]
	s.index--
}

func (s *stack) empty() bool {
	return s.index == 0
}

func (s *stack) top() int {
	return s.list[s.index-1]
}

func (s *stack) push(e int) {
	if s.index > len(s.list)-1 {
		s.list = append(s.list, e)
	} else {
		s.list[s.index] = e
	}
	s.index++
}

func trap(height []int) int {
	s := newStack()
	var result int
	min := func(a, b int) int {
		m := a
		if b < a {
			m = b
		}
		return m
	}
	for i := range height {
		for !s.empty() && height[s.top()] < height[i] {
			top := s.top()
			s.pop()
			if s.empty() {
				break
			}
			depth := min(height[s.top()], height[i]) - height[top]
			width := i - s.top() - 1
			result += width * depth
		}
		s.push(i)
	}
	return result
}
