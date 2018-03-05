package binary_search_tree

const (
	red   color = true
	black color = false
)

// Restrict of RedBlackTree:
// 1. the color of tree node can be red or black
// 2. the root is black
// 3. each leaf(nil node) is black
// 4. both left and right children of each red node must be black
// 5. all the routes from any node to all its leaf children contain the same number of black nodes

type RedBlackTree struct {
	node *treeNode
}

func (tree *RedBlackTree) Search(value int) bool {
	_, target := findFromNode(tree.node, value)
	return target != nil
}

func (tree *RedBlackTree) Insert(value int) {
	parent, current := findFromNode(tree.node, value)
	if current != nil {
		return
	}
	current = &treeNode{color: red, value: value}
	if parent.value > value {
		parent.left = current
	} else {
		parent.right = current
	}
	insertFixUp(current)
	tree.node.color = black
}

func leftRotateRedBlack(node *treeNode) *treeNode {
	newNode := node.right
	newNode.parent = node.parent
	node.right = newNode.left
	node.right.parent = node
	newNode.left = node
	node.parent = newNode
	if node == newNode.parent.left {
		newNode.parent.left = newNode
	} else {
		newNode.parent.right = newNode
	}
	return newNode
}

func rightRotateRedBlack(node *treeNode) *treeNode {
	newNode := node.left
	newNode.parent = node.parent
	node.left = newNode.right
	node.left.parent = node
	newNode.right = node
	node.parent = newNode
	if node == newNode.parent.left {
		newNode.parent.left = newNode
	} else {
		newNode.parent.right = newNode
	}
	return newNode
}

func insertFixUp(node *treeNode) {
	if node == nil {
		return
	}
	var parent, gparent *treeNode = node.parent, nil
	// if parent exists and its color is red
	for parent != nil && parent.color == red {
		gparent = parent.parent
		// if parent if the left child of grandparent
		if gparent.left == parent {
			uncle := gparent.right
			// if uncle exists and its color is red
			if uncle != nil && uncle.color == red {
				parent.color = black
				uncle.color = black
				gparent.color = red
				node = gparent
				continue
			}
			// if the node is the right child of parent and uncle's color is black
			if parent.right == node {
				parent = leftRotateRedBlack(parent)
				node = parent.left
			}
			// if the node is the left child of parent and uncle's color is black
			parent.color = black
			gparent.color = red
			gparent = rightRotateRedBlack(gparent)
		} else {
			uncle := gparent.left
			// if uncle exists and its color is red
			if uncle != nil && uncle.color == red {
				parent.color = black
				uncle.color = black
				gparent.color = red
				node = gparent
				continue
			}
			// if the node is the left child of parent and uncle's color is black
			if parent.left == node {
				parent = rightRotateRedBlack(parent)
				node = parent.right
			}
			// if the node is the right child of parent and uncle's color is black
			parent.color = black
			gparent.color = red
			gparent = leftRotateRedBlack(gparent)
		}
	}

}

func (tree *RedBlackTree) Delete(value int) {
	_, target := findFromNode(tree.node, value)
	if target != nil {
		deleteNodeFromBRedBlackTree(tree, target)
	}
}

func deleteNodeFromBRedBlackTree(tree *RedBlackTree, node *treeNode) {
	var (
		child, parent *treeNode
		color         color
	)
	parent = node.parent
	if node.left == nil {
		if parent.left == node {
			parent.left = node.right
			node.parent = nil
		} else {
			parent.right = node.right
			node.parent = nil
		}
	} else if node.right == nil {
		if parent.left == node {
			parent.left = node.left
			node.parent = nil
		} else {
			parent.right = node.left
			node.parent = nil
		}
	} else {
		// two children are not nil
		// 1. find out replace
		replace := node.left
		for replace.left != nil {
			replace = replace.left
		}
		// 2.
		if parent != nil {
			if parent.left == node {
				parent.left = replace
			} else {
				parent.right = replace
			}
		} else {
			tree.node = replace
		}
		child, parent, color = replace.right, replace.parent, replace.color
		if parent == node {
			parent = replace
		} else {
			if child != nil {
				child.parent = parent
				parent.left = child
			}
			replace.right = node.right
			replace.right.parent = replace
		}
		replace.parent = node.parent
		replace.color = node.color
		replace.left = node.left
		replace.left.parent = replace
		if color == black {
			// fix up
			deleteFixUp(child, parent, tree.node)
		}
	}
}

// the color of node must be black
func deleteFixUp(node, parent, root *treeNode) {
	var other *treeNode
	for (node == nil || node.color == black) && node != root {
		if parent.left == node {
			other = parent.right
			// 1. other is red
			if other.color == red {
				other.color = black
				parent.color = red
				leftRotateRedBlack(parent)
				other = parent.right
			}
			// 2. other is black
			if (other.left == nil || other.left.color == black) && (other.right == nil || other.right.color == black) {
				// both children of other are black
				other.color = red
				node = parent
				parent = node.parent
			} else {
				// the left child of other is red and the right one is black
				if other.right == nil || other.right.color == black {
					//
					other.left.color = black
					other.color = red
					rightRotateRedBlack(other)
					other = parent.right
				}
				// the right child of other is red
				other.color = parent.color
				parent.color = black
				other.right.color = black
				leftRotateRedBlack(parent)
				other = root
				break
			}
		} else {
			other = parent.left
			// 1. other is red
			if other.color == red {
				other.color = black
				parent.color = red
				rightRotateRedBlack(parent)
				other = parent.left
			}
			// 2. other is black
			if (other.left == nil || other.left.color == black) && (other.right == nil || other.right.color == black) {
				// both children of other are black
				other.color = red
				node = parent
				parent = node.parent
			} else {
				// the right child of other is red and the left one is black
				if other.left == nil || other.left.color == black {
					//
					other.right.color = black
					other.color = red
					leftRotateRedBlack(other)
					other = parent.left
				}
				// the left child of other is red
				other.color = parent.color
				parent.color = black
				other.left.color = black
				rightRotateRedBlack(parent)
				other = root
				break
			}
		}
	}
	if node != nil {
		node.color = black
	}
}
