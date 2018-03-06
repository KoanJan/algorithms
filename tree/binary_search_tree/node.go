package binary_search_tree

import "fmt"

type color bool

type treeNode struct {
	color               color
	height              int
	value               int
	left, right, parent *treeNode
}

func findFromNode(node *treeNode, value int) (parent, target *treeNode) {

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

func formatNode(node *treeNode) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d-(%s, %s)", node.value, formatNode(node.left), formatNode(node.right))
}

func leftRotate(node *treeNode) {
	parent := node.parent
	node.parent = node.right
	node.parent.parent = parent
	node.right = node.parent.left
	node.parent.left = node
	if parent == nil {
		if parent.left == node {
			parent.left = node.parent
		} else {
			parent.right = node.parent
		}
	}
}

func rightRotate(node *treeNode) {
	parent := node.parent
	node.parent = node.left
	node.parent.parent = parent
	node.left = node.parent.right
	node.parent.right = node
	if parent == nil {
		if parent.left == node {
			parent.left = node.parent
		} else {
			parent.right = node.parent
		}
	}
}
