// Time        : 2019/01/24
// Description :

package valid_parentheses

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.

type Stack struct {
	list []interface{}
}

func (s *Stack) Push(e interface{}) {
	s.list = append(s.list, e)
}

func (s *Stack) Pop() interface{} {
	if len(s.list) == 0 {
		panic("no elem")
	}
	elem := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return elem
}

func (s *Stack) Count() int {
	return len(s.list)
}

func isValid(s string) bool {
	var matched = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stk := &Stack{}
	ss := []rune(s)
	for _, r := range ss {
		if r == '(' || r == '{' || r == '[' {
			stk.Push(r)
		} else {
			if stk.Count() == 0 {
				return false
			}
			if stk.Pop().(rune) != matched[r] {
				return false
			}
		}
	}
	if stk.Count() != 0 {
		return false
	}
	return true
}
