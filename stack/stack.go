package stack

import "errors"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return -1, errors.New("Stack is empty! Nothing to pop")
	}

	stackSize := len(s.items) - 1

	topItem := s.items[stackSize]

	s.items = s.items[:stackSize]

	return topItem, nil
}
