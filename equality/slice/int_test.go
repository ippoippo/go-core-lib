package slice

import (
	"testing"
)

func TestEqualInt(t *testing.T) {
	tests := []struct {
		a []int
		b []int
		expected bool
		desc     string // Description of testcase
	}{
		{nil, nil, true, "Two nils"},
		{[]int{}, nil, false, "b is nil"},
		{nil, []int{}, false, "a is nil"},
		{[]int{}, []int{}, true, "Two empty"},
		{[]int{1}, []int{1}, true, "Two 1 element"},
		{[]int{1, 2, 4}, []int{1, 2, 4}, true, "Two 3 element"},
		{[]int{1, 2, 4}, []int{1, 2}, false, "Mistmatch element count"},
		{[]int{1, 2, 4}, []int{1, 4, 2}, false, "Mistmatch elements"},
	}

	for _, test := range tests {
		if actual := EqualInt(test.a, test.b); actual != test.expected {
			t.Errorf("[%s]:String() = Actual:[%v], Expected:[%v]", test.desc, actual, test.expected)
		}
	}
}

