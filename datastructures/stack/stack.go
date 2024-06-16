package stack

import "errors"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var zeroValue T
	size := len(s.items)
	if size > 0 {
		poppedItem := s.items[size-1]
		s.items = s.items[:(size - 1)]
		return poppedItem, nil
	} else {
		return zeroValue, errors.New("error - stack is empty")
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Clear() {
	clear(s.items)
}

func (s *Stack[T]) Top() (T, error) {
	var zeroValue T
	size := len(s.items)
	if size > 0 {
		return s.items[size-1], nil
	} else {
		return zeroValue, errors.New("error - stack is empty")
	}
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
