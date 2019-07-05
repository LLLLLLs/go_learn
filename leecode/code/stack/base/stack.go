// Time        : 2019/01/24
// Description :

package base

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
