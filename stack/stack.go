package stack

import "errors"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("Stack is empty! Nothing to pop")
	}

	stackSize := len(s.items) - 1

	topItem := s.items[stackSize]

	s.items = s.items[:stackSize]

	return topItem, nil
}
