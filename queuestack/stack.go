package queuestack

import "errors"

// Stack represents a collection data structure with LIFO (Last In First Out) semantics).
type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Empty() bool
}

type stack struct {
	data []interface{}
	size int
}

func (s *stack) Push(value interface{}) {
	s.data = append(s.data, value)
	s.size++
}

func (s *stack) Pop() (interface{}, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		s.size--
		s.data = s.data[:s.size]
		return value, nil
	}

	return nil, errors.New("no such element")
}

func (s *stack) Peek() (interface{}, error) {
	if s.size > 0 {
		value := s.data[s.size-1]
		return value, nil
	}

	return nil, errors.New("no such element")
}

func (s *stack) Empty() bool {
	return s.size == 0
}

// NewStack returns an empty stack.
func NewStack() Stack {
	return &stack{}
}
