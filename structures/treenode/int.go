package treenode

import (
	"fmt"
)

// A Int is a binary tree node, storing an `int` value, and optional `string` Label
type Int struct {
	Left  *Int
	Value int
	Label string
	Right *Int
}

// NewIntRoot returns a new node, with no children set
func NewIntRoot(val int, lbl string) *Int {
	return &Int{
		Value: val,
		Label: lbl,
	}
}

// NewIntRootWithChildren returns a new node, with children set
func NewIntRootWithChildren(val int, lbl string, l *Int, r *Int) *Int {
	return &Int{
		Value: val,
		Label: lbl,
		Left:  l,
		Right: r,
	}
}

// String returns full dump of the contents of the node
func (i *Int) String() string {
	if i == nil {
		return "()"
	}
	s := ""

	if i.Left != nil {
		s += i.Left.String() + " "
	}

	s += fmt.Sprint(i.Value)
	if i.Label != "" {
		s += fmt.Sprintf("[%s]", i.Label)
	}

	if i.Right != nil {
		s += " " + i.Right.String()
	}
	return "(" + s + ")"
}

// Desc returns the `Label` if it exists, otherwise it returns a string representation of the `Value`
func (i *Int) Desc() string {
	if i == nil {
		return ""
	}
	if i.Label == "" {
		return fmt.Sprintf("%d", i.Value)
	}
	return i.Label
}
