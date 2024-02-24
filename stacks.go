package main

import "errors"

type Stack struct {
	elements []any
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(el any) {
	s.elements = append(s.elements, el)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	last := len(s.elements) - 1
	el := s.elements[last]
	s.elements = s.elements[0:last]
	return el, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}
