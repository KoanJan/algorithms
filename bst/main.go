package main

import (
	"fmt"
)

// test
func main() {

	bst := &BST{Val: POI(25)}
	fmt.Printf("INIT: %s\n", bst.String())

	bst.Insert(9)
	bst.Insert(50)
	bst.Insert(16)
	bst.Insert(30)
	bst.Insert(2)
	bst.Insert(65)
	fmt.Printf("AFTER INSERT: \n%s\n", bst.String())

	bst.Delete(9)
	fmt.Printf("AFTER DELETE: \n%s\n", bst.String())

	bst.Delete(25)
	fmt.Printf("AFTER DELETE: \n%s\n", bst.String())
}
