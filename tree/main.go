package main

import (
	"fmt"

	"algorithms/tree/binary_search_tree"
)

func main() {
	tree := binary_search_tree.NewAVLTree()
	fmt.Println(tree)
	tree.Insert(5)
	fmt.Println(tree)
	tree.Insert(3)
	fmt.Println(tree)
	tree.Insert(7)
	fmt.Println(tree)
	tree.Insert(2)
	fmt.Println(tree)
	tree.Insert(4)
	fmt.Println(tree)
	tree.Insert(6)
	fmt.Println(tree)
	tree.Insert(8)
	fmt.Println(tree)

	target := 2
	fmt.Printf("find %d: %t\n", target, tree.Search(target))

	target = 3
	fmt.Printf("find %d: %t\n", target, tree.Search(target))

	tree.Delete(target)
	fmt.Println(tree)

	target = 4
	fmt.Printf("find %d: %t\n", target, tree.Search(target))

	tree.Delete(target)
	fmt.Println(tree)

	target = 2
	fmt.Printf("find %d: %t\n", target, tree.Search(target))

	tree.Delete(target)
	fmt.Println(tree)

	fmt.Printf("find %d: %t\n", target, tree.Search(target))

}
