package binary_search_tree

import "fmt"

type AVLTree struct {
	node *treeNode
}

func (tree *AVLTree) Search(value int) bool {
	_, target := findFromNode(tree.node, value)
	return target != nil
}

func (tree *AVLTree) Insert(value int) {
	tree.node = insertNodeIntoAVLTree(tree.node, value)
}

func insertNodeIntoAVLTree(node *treeNode, value int) *treeNode {
	if node == nil {
		node = &treeNode{height: 1, value: value}
		return node
	}
	if node.value > value {
		node.left = insertNodeIntoAVLTree(node.left, value)
		fixHeight(node)
		if height(node.left)-height(node.right) > 1 {
			if height(node.left.left) > height(node.left.right) {
				// right rotate
				node = rightRotateAVL(node)
			} else {
				// left right rotate
				node.right = leftRotateAVL(node.right)
				node = rightRotateAVL(node)
			}
		}
	} else if node.value < value {
		node.right = insertNodeIntoAVLTree(node.right, value)
		fixHeight(node)
		if height(node.right)-height(node.left) > 1 {
			if height(node.right.right) > height(node.right.left) {
				// left rotate
				node = leftRotateAVL(node)
			} else {
				// right left rotate
				node.left = rightRotateAVL(node.left)
				node = leftRotateAVL(node)
			}
		}
	}
	return node
}

func (tree *AVLTree) Delete(value int) {
	tree.node = deleteNodeFromAVLTree(tree.node, value)
}

func deleteNodeFromAVLTree(node *treeNode, value int) *treeNode {

	if node != nil {
		if node.value == value {
			if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else if height(node.left) > height(node.right) {
				newValue := findMax(node.left).value
				node.value = newValue
				node.left = deleteNodeFromAVLTree(node.left, newValue)
			} else {
				newValue := findMin(node.right).value
				node.value = newValue
				node.right = deleteNodeFromAVLTree(node.right, newValue)
			}
		} else if node.value > value {
			node.left = deleteNodeFromAVLTree(node.left, value)
			if height(node.right)-height(node.left) > 1 {
				if height(node.right.left) > height(node.right.right) {
					node.right = rightRotateAVL(node.right)
				}
				node = leftRotateAVL(node)
			}
		} else {
			node.right = deleteNodeFromAVLTree(node.right, value)
			if height(node.left)-height(node.right) > 1 {
				if height(node.left.right) > height(node.left.left) {
					node.left = leftRotateAVL(node.left)
				}
				node = rightRotateAVL(node)
			}
		}
	}
	return node
}

func findMax(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	rightChild := node.right
	if rightChild == nil {
		return node
	}
	for rightChild.right != nil {
		rightChild = rightChild.right
	}
	return rightChild
}

func findMin(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	leftChild := node.left
	if leftChild == nil {
		return node
	}
	for leftChild.left != nil {
		leftChild = leftChild.left
	}
	return leftChild
}

func rightRotateAVL(node *treeNode) *treeNode {
	newNode := node.left
	node.left = newNode.right
	newNode.right = node
	node = newNode
	// fix height
	fixHeight(node.right)
	fixHeight(node)
	return node
}

func leftRotateAVL(node *treeNode) *treeNode {
	newNode := node.right
	node.right = newNode.left
	newNode.left = node
	node = newNode
	// fix height
	fixHeight(node.left)
	fixHeight(node)
	return node
}

func fixHeight(node *treeNode) {
	if node == nil {
		return
	}
	node.height = max(height(node.left), height(node.right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func height(node *treeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (tree *AVLTree) String() string {
	return fmt.Sprintf("AVLTree: {node: %s}", formatNode(tree.node))
}

func NewAVLTree() *AVLTree {
	tree := new(AVLTree)
	return tree
}
