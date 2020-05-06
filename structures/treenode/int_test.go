package treenode

import (
	"testing"
)

func TestString(t *testing.T) {
	var nilNode *Int
	single := NewIntRoot(1, "One")
	left := NewIntRoot(1, "One")
	right := NewIntRoot(3, "Three")
	root := NewIntRootWithChildren(2, "Two", left, right)
	nolabel := NewIntRoot(10, "")

	tests := []struct {
		node     *Int   // Struct test executes on
		expected string // Expected result
		desc     string // Description of testcase
	}{
		{nilNode, "()", "Nil node"},
		{single, "(1[One])", "Single with label"},
		{nolabel, "(10)", "Single with no label"},
		{root, "((1[One]) 2[Two] (3[Three]))", "Root with two children"},
	}

	for _, test := range tests {
		if actual := test.node.String(); actual != test.expected {
			t.Errorf("[%s]:String() = Actual:[%q], Expected:[%q]", test.desc, actual, test.expected)
		}
	}
}

func TestDesc(t *testing.T) {
	var nilNode *Int
	single := NewIntRoot(1, "One")
	left := NewIntRoot(1, "One")
	right := NewIntRoot(3, "Three")
	root := NewIntRootWithChildren(2, "Two", left, right)
	nolabel := NewIntRoot(10, "")

	tests := []struct {
		node     *Int   // Struct test executes on
		expected string // Expected result
		desc     string // Description of testcase
	}{
		{nilNode, "", "Nil node"},
		{single, "One", "Single with label"},
		{nolabel, "10", "Single with no label"},
		{root, "Two", "Root with two children"},
	}

	for _, test := range tests {
		if actual := test.node.Desc(); actual != test.expected {
			t.Errorf("[%s]:Desc() = Actual:[%q], Expected:[%q]", test.desc, actual, test.expected)
		}
	}
}

func TestEqual(t *testing.T) {
	var nilNode *Int
	s1 := NewIntRoot(1, "")
	s2 := NewIntRoot(2, "")
	s1b := NewIntRoot(1, "")

	tests := []struct {
		node     *Int   // Struct test executes on
		input *Int // Parameter to the function
		expected bool // Expected result
		desc     string // Description of testcase
	}{
		{nilNode, nilNode, true, "Nil nodes"},
		{nilNode, s1, false, "operating on nil"},
		{s1, nilNode, false, "nil parameter"},
		{s1, s1, true, "Itself"},
		{s1, s1b, true, "Different, but same value"},
		{s1, s2, false, "Different value"},
	}

	for _, test := range tests {
		if actual := test.node.Equal(test.input); actual != test.expected {
			t.Errorf("[%s]:Equal() = Actual:[%v], Expected:[%v]", test.desc, actual, test.expected)
		}
	}
}

func TestLessThan(t *testing.T) {
	var nilNode *Int
	s1 := NewIntRoot(1, "")
	s2 := NewIntRoot(2, "")
	s1b := NewIntRoot(1, "")

	tests := []struct {
		node     *Int   // Struct test executes on
		input *Int // Parameter to the function
		expected bool // Expected result
		desc     string // Description of testcase
	}{
		{nilNode, nilNode, false, "Nil nodes"},
		{nilNode, s1, false, "operating on nil"},
		{s1, nilNode, false, "nil parameter"},
		{s1, s1, false, "Itself"},
		{s1, s1b, false, "Different, but same value"},
		{s1, s2, true, "parameter is more"},
		{s2, s1, false, "parameter is less"},
	}

	for _, test := range tests {
		if actual := test.node.LessThan(test.input); actual != test.expected {
			t.Errorf("[%s]:LessThan() = Actual:[%v], Expected:[%v]", test.desc, actual, test.expected)
		}
	}
}

func TestMoreThan(t *testing.T) {
	var nilNode *Int
	s1 := NewIntRoot(1, "")
	s2 := NewIntRoot(2, "")
	s1b := NewIntRoot(1, "")

	tests := []struct {
		node     *Int   // Struct test executes on
		input *Int // Parameter to the function
		expected bool // Expected result
		desc     string // Description of testcase
	}{
		{nilNode, nilNode, false, "Nil nodes"},
		{nilNode, s1, false, "operating on nil"},
		{s1, nilNode, false, "nil parameter"},
		{s1, s1, false, "Itself"},
		{s1, s1b, false, "Different, but same value"},
		{s1, s2, false, "parameter is more"},
		{s2, s1, true, "parameter is less"},
	}

	for _, test := range tests {
		if actual := test.node.MoreThan(test.input); actual != test.expected {
			t.Errorf("[%s]:MoreThan() = Actual:[%v], Expected:[%v]", test.desc, actual, test.expected)
		}
	}
}
