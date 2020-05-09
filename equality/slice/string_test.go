package slice

import (
	"testing"
)

func TestEqualString(t *testing.T) {
	tests := []struct {
		a        []string
		b        []string
		expected bool
		desc     string // Description of testcase
	}{
		{nil, nil, true, "Two nils"},
		{[]string{}, nil, false, "b is nil"},
		{nil, []string{}, false, "a is nil"},
		{[]string{}, []string{}, true, "Two empty"},
		{[]string{"1"}, []string{"1"}, true, "Two 1 element"},
		{[]string{"1", "2", "4"}, []string{"1", "2", "4"}, true, "Two 3 element"},
		{[]string{"1", "2", "4"}, []string{"1", "2"}, false, "Mistmatch element count"},
		{[]string{"1", "2", "4"}, []string{"1", "4", "2"}, false, "Mistmatch elements"},
	}

	for _, test := range tests {
		if actual := EqualString(test.a, test.b); actual != test.expected {
			t.Errorf("[%s]:EqualString() = Actual:[%v], Expected:[%v]", test.desc, actual, test.expected)
		}
	}
}
