package stack

import (
	"errors"
)

// A String is a simple slice based stack of `string` values
type String struct {
	Stack []string
}

// NewStringStack creates the stack
func NewStringStack() *String {
	return &String{
		Stack: []string{},
	}
}

// Push a val onto the stack
func (s *String) Push(val string) *String {
	s.Stack = append(s.Stack, val)
	return s
}

// Pop topmost item off stack, return it and the remaining stack
func (s *String) Pop() (string, *String, error) {
	l := len(s.Stack)
	if l == 0 {
		// Error case
		return "", s, errors.New("Cannot pop from empty stack")
	}
	result := s.Stack[l-1]
	s.Stack = s.Stack[:l-1]
	return result, s, nil
}

// Count returns item count of the stack
func (s *String) Count() int {
	return len(s.Stack)
}

// AsSlice returns the contents as an array slice
func (s *String) AsSlice() []string {
	return s.Stack
}
