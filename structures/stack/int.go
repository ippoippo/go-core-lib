package stack

import (
	"errors"
)

// A Int is a simple slice based stack of `int` values
type Int struct {
	Stack []int
}

// NewIntStack creates the stack
func NewIntStack() *Int {
	return &Int{
		Stack: []int{},
	}
}

// Push a val onto the stack
func (s *Int) Push(val int) *Int {
	s.Stack = append(s.Stack, val)
	return s
}

// Pop topmost item off stack, return it and the remaining stack
func (s *Int) Pop() (int, *Int, error) {
	l := len(s.Stack)
	if l == 0 {
		// Error case
		return 0, s, errors.New("Cannot pop from empty stack")
	}
	result := s.Stack[l-1]
	s.Stack = s.Stack[:l-1]
	return result, s, nil
}

// Count returns item count of the stack
func (s *Int) Count() int {
	return len(s.Stack)
}

// AsSlice returns the contents as an array slice
func (s *Int) AsSlice() []int {
	return s.Stack
}
