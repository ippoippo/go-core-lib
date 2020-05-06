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
