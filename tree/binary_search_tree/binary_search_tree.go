package binary_search_tree

import (
	"fmt"
)

type BinarySearchTree struct {
	node *node
}

type node struct {
	value       int
	left, right *node
}

func findFromNode(node *node, value int) (parent, target *node) {

	target = node
	for target != nil && target.value != value {
		parent = target
		if target.value > value {
			target = target.left
		} else {
			target = target.right
		}
	}
	return parent, target
}

func (bst *BinarySearchTree) Search(value int) bool {
	_, target := findFromNode(bst.node, value)
	return target != nil
}

func (bst *BinarySearchTree) Insert(value int) {
	if bst.node == nil {
		bst.node = &node{value, nil, nil}
		return
	}
	parent, target := findFromNode(bst.node, value)
	if target != nil {
		return
	}
	if parent.value > value {
		parent.left = &node{value, nil, nil}
	} else {
		parent.right = &node{value, nil, nil}
	}
}

func (bst *BinarySearchTree) Delete(value int) {
	if bst.node == nil {
		return
	}
	parent, target := findFromNode(bst.node, value)
	if target == nil {
		return
	}
	if parent == nil {
		bst.node = nil
		return
	}
	isLeftChild := parent.value > target.value

	// select the node to swap and construct a new node
	var newNode *node = nil
	if target.right != nil && target.left != nil {
		oldParentOfNewNode := target
		newNode = target.right
		for newNode.left != nil {
			oldParentOfNewNode = newNode
			newNode = newNode.left
		}
		if newNode == target.right {
			newNode.left = target.left
		} else {
			newLeaf := newNode.right
			newNode.right = target.right
			newNode.left = target.left
			oldParentOfNewNode.left = newLeaf
		}
	} else if target.right != nil {
		newNode = target.right
	} else if target.left != nil {
		newNode = target.left
	}

	// swap
	if isLeftChild {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (bst *BinarySearchTree) String() string {
	return fmt.Sprintf("BinarySearchTree: {node: %s}", formatNode(bst.node))
}

func formatNode(node *node) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d-(%s, %s)", node.value, formatNode(node.left), formatNode(node.right))
}
