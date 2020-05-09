package stack

import (
	"fmt"
	"testing"

	"github.com/ippoippo/go-core-lib/equality/slice"
)

func TestNewStringStack(t *testing.T) {
	s := NewStringStack()
	//Check Count is zero
	if s.Count() != 0 {
		t.Errorf("NewStringStack:[%v]", s)
	}
}

func TestStringPush(t *testing.T) {
	tests := []struct {
		values   []string
		expected string
	}{
		{
			values:   []string{"1"},
			expected: "&{[1]}",
		},
		{
			values:   []string{"1", "3", "5"},
			expected: "&{[1 3 5]}",
		},
		{
			values:   []string{"1", "-10", "5", "4"},
			expected: "&{[1 -10 5 4]}",
		},
	}

	for _, test := range tests {
		s := NewStringStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		if actual := fmt.Sprintf("%v", s); actual != test.expected {
			t.Errorf("[%v]:Push() = Actual:[%v], Expected:[%v]", test.values, actual, test.expected)
		}
	}
}

func TestStringCount(t *testing.T) {
	tests := []struct {
		values   []string
		expected int
	}{
		{
			values:   []string{"1"},
			expected: 1,
		},
		{
			values:   []string{"1", "3", "5"},
			expected: 3,
		},
		{
			values:   []string{"1", "-10", "5", "4"},
			expected: 4,
		},
	}

	for _, test := range tests {
		s := NewStringStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		if actual := s.Count(); actual != test.expected {
			t.Errorf("[%v]:Count() = Actual:[%v], Expected:[%v]", test.values, actual, test.expected)
		}
	}
}

func TestStringPop(t *testing.T) {
	tests := []struct {
		values      []string
		expectedV   string
		expectedS   string
		expectedErr bool
	}{
		{
			values:      []string{"1"},
			expectedV:   "1",
			expectedS:   "&{[]}",
			expectedErr: false,
		},
		{
			values:      []string{"1", "3", "5"},
			expectedV:   "5",
			expectedS:   "&{[1 3]}",
			expectedErr: false,
		},
		{
			values:      []string{},
			expectedV:   "0",
			expectedS:   "",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		s := NewStringStack()
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

func TestStringAsSlice(t *testing.T) {
	tests := []struct {
		values []string
	}{
		{
			values: []string{"1"},
		},
		{
			values: []string{"1", "3", "5"},
		},
		{
			values: []string{"1", "-10", "5", "4"},
		},
	}

	for _, test := range tests {
		s := NewStringStack()
		for _, v := range test.values {
			s = s.Push(v)
		}

		actual := s.AsSlice()
		if !slice.EqualString(actual, test.values) {
			t.Errorf("[%v]:AsSlice() = Actual:[%v], Expected:[%v]", test.values, actual, test.values)
		}
	}
}
