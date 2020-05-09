package stack

import (
	"fmt"
	"testing"

	"github.com/ippoippo/go-core-lib/equality/slice"
)

func TestNewIntStack(t *testing.T) {
	s := NewIntStack()
	//Check Count is zero
	if s.Count() != 0 {
		t.Errorf("NewIntStack:[%v]", s)
	}
}

func TestIntPush(t *testing.T) {
	tests := []struct {
		values   []int
		expected string
	}{
		{
			values:   []int{1},
			expected: "&{[1]}",
		},
		{
			values:   []int{1, 3, 5},
			expected: "&{[1 3 5]}",
		},
		{
			values:   []int{1, -10, 5, 4},
			expected: "&{[1 -10 5 4]}",
		},
	}

	for _, test := range tests {
		s := NewIntStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		if actual := fmt.Sprintf("%v", s); actual != test.expected {
			t.Errorf("[%v]:Push() = Actual:[%v], Expected:[%v]", test.values, actual, test.expected)
		}
	}
}

func TestIntCount(t *testing.T) {
	tests := []struct {
		values   []int
		expected int
	}{
		{
			values:   []int{1},
			expected: 1,
		},
		{
			values:   []int{1, 3, 5},
			expected: 3,
		},
		{
			values:   []int{1, -10, 5, 4},
			expected: 4,
		},
	}

	for _, test := range tests {
		s := NewIntStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		if actual := s.Count(); actual != test.expected {
			t.Errorf("[%v]:Count() = Actual:[%v], Expected:[%v]", test.values, actual, test.expected)
		}
	}
}

func TestIntPop(t *testing.T) {
	tests := []struct {
		values      []int
		expectedV   int
		expectedS   string
		expectedErr bool
	}{
		{
			values:      []int{1},
			expectedV:   1,
			expectedS:   "&{[]}",
			expectedErr: false,
		},
		{
			values:      []int{1, 3, 5},
			expectedV:   5,
			expectedS:   "&{[1 3]}",
			expectedErr: false,
		},
		{
			values:      []int{},
			expectedV:   0,
			expectedS:   "",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		s := NewIntStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		actualV, actualS, err := s.Pop()
		if test.expectedErr {
			if err == nil {
				t.Errorf("[%v]:Pop() = Expected err not received", test.values)
			}
		} else {
			if err != nil {
				t.Errorf("[%v]:Pop() = Unexpected err [%v]", test.values, err)
			}
			if actualV != test.expectedV {
				t.Errorf("[%v]:Pop() Value = Actual:[%v], Expected:[%v]", test.values, actualV, test.expectedV)
			}
			if fmt.Sprintf("%v", actualS) != test.expectedS {
				t.Errorf("[%v]:Pop() Remaining Stack = Actual:[%v], Expected:[%v]", test.values, actualV, test.expectedV)
			}
		}
	}
}

func TestIntAsSlice(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{
			values: []int{1},
		},
		{
			values: []int{1, 3, 5},
		},
		{
			values: []int{1, -10, 5, 4},
		},
	}

	for _, test := range tests {
		s := NewIntStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		actual := s.AsSlice()
		if !slice.EqualInt(actual, test.values) {
			t.Errorf("[%v]:AsSlice() = Actual:[%v], Expected:[%v]", test.values, actual, test.values)
		}
	}
}
