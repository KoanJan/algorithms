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
	fixUp(current)
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

func fixUp(node *treeNode) {
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
	// TODO
}
