package main

import (
	"fmt"
)

func main() {

	//test
	tree := NewBNode(1, 12)
	fmt.Printf("tree: %v\n\n", tree)

	tree.Insert(2)
	fmt.Printf("tree: %v\n\n", tree)

	tree.Insert(4)
	fmt.Printf("tree: %v\n\n", tree)

	tree.Insert(5)
	fmt.Printf("tree: %v\n\n", tree)

	var (
		k     int
		node  *BNode
		index int
	)

	k = 5
	node, index = tree.Search(k)
	fmt.Printf("search %d in tree, node is %v and index is %d\n", k, node, index)

	k = 2
	node, index = tree.Search(k)
	fmt.Printf("search %d in tree, node is %v and index is %d\n", k, node, index)

	k = 4
	node, index = tree.Search(k)
	fmt.Printf("search %d in tree, node is %v and index is %d\n", k, node, index)
}
