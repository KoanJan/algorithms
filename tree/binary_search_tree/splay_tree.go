package binary_search_tree

type SplayTree struct {
	node *treeNode
}

func (tree *SplayTree) Search(value int) bool {
	_, target := findFromNode(tree.node, value)
	if target == nil {
		return false
	}
	tree.node = splay(target)
	return true
}

func (tree *SplayTree) Insert(value int) {
	tree.node = insertNodeIntoBST(tree.node, value)
}

func (tree *SplayTree) Delete(value int) {
	if tree.Search(value) {
		node := tree.node
		if node.right != nil {
			min := findMin(node.right)
			if min.parent == node {
				node.value = min.value
				node.right = min.right
			} else {
				node.value = min.value
				min.parent.left = min.right
				min.parent = nil
			}
		} else {
			newRoot := node.left
			newRoot.parent = nil
			node.left = nil
			tree.node = newRoot
		}
	}
}

func splay(node *treeNode) *treeNode {
	for node != nil && node.parent != nil && node.parent.parent != nil {
		node = upInSplayNode(node)
	}
	if node == nil {
		return node
	}
	parent := node.parent
	if parent == nil {
		return node
	}
	gparent := parent.parent
	if gparent == nil {
		if parent.left == node {
			rightRotate(parent)
		} else {
			leftRotate(parent)
		}
	}
	return node
}

func upInSplayNode(node *treeNode) *treeNode {
	parent := node.parent
	gparent := parent.parent
	left1 := parent.left == node
	left2 := gparent.left == parent
	if left1 && left2 {
		rightRotate(parent)
		rightRotate(gparent)
	} else if left1 {
		rightRotate(parent)
		leftRotate(gparent)
	} else if left2 {
		leftRotate(parent)
		rightRotate(gparent)
	} else {
		leftRotate(parent)
		leftRotate(gparent)
	}
	return node
}
