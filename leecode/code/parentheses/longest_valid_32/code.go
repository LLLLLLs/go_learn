// Time        : 2019/06/28
// Description :

package longest_valid_32

// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
// Example 1:
//
// Input: "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()"
// Example 2:
//
// Input: ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()"

type stack struct {
	list  []int
	index int
}

func newStack() *stack {
	return &stack{
		list: make([]int, 0),
	}
}

func (s *stack) push(i int) {
	if s.index < len(s.list) {
		s.list[s.index] = i
	} else {
		s.list = append(s.list, i)
	}
	s.index++
}

func (s *stack) top() int {
	return s.list[len(s.list)-1]
}

func (s *stack) pop() int {
	num := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	s.index--
	return num
}

func (s *stack) isEmpty() bool {
	return len(s.list) == 0
}

func longestValidParentheses(s string) int {
	stack := newStack()
	stack.push(-1)
	result := 0
	for i := range s {
		if s[i] == '(' {
			stack.push(i)
		} else {
			stack.pop()
			if stack.isEmpty() {
				stack.push(i)
			} else {
				index := stack.top()
				if result < i-index {
					result = i - index
				}
			}
		}
	}
	return result
}
