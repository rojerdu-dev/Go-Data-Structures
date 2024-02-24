package main

import (
	"fmt"
	"testing"
)

func TestStackInt(t *testing.T) {
	s := NewStack()

	s.Push(11)
	s.Push(23)
	s.Push(37)

	fmt.Printf("Current Stack: %v \n", s.elements)

	testCases := []struct {
		expected int
	}{
		{37},
		{23},
		{11},
	}

	for _, tc := range testCases {
		el, err := s.Pop()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if el != tc.expected {
			t.Errorf("expected %v, got %v", tc.expected, el)
		}
	}

	if !s.IsEmpty() {
		t.Errorf("expected empty stack, got %v", s.elements)
	}
}
