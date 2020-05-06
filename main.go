package main

import (
	"fmt"
	"github.com/ippoippo/go-core-lib/structures/treenode"
)

func main() {
	tn := treenode.NewIntRoot(1, "One")
	fmt.Printf("%v\n", tn.String())
}