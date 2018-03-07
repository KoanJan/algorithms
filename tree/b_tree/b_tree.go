package b_tree

type BTree struct {
	node *treeNode
	t    int
}

func (tree *BTree) Search(key int) bool {
	_, exist := findInBTreeNode(tree.node, key)
	return exist
}

func findInBTreeNode(node *treeNode, key int) (*treeNode, bool) {
	childNode := node
	for childNode != nil {
		idx, exist := findInKeys(childNode.keys, key)
		if exist {
			return childNode, true
		}
		if childNode.leaf {
			break
		}
		childNode = childNode.children[idx]
	}
	return childNode, false
}

func (tree *BTree) Insert(key int) {
	node, exist := findInBTreeNode(tree.node, key)
	if exist {
		return
	}
	insertKeyIntoLeafNode(node, key)
	tree.node = split(node, tree.t)
}

func split(node *treeNode, t int) *treeNode {
	// not necessary to split
	if len(node.keys) <= 2*t-1 {
		for node.parent != nil {
			node = node.parent
		}
		return node
	}
	// split
	upKey, leftKeys, rightKeys := node.keys[t], node.keys[0:t], node.keys[t+1:]
	var leftChildren, rightChildren []*treeNode
	if !node.leaf {
		leftChildren = node.children[0 : t+1]
		rightChildren = node.children[t+1:]
	}
	left := &treeNode{leftKeys, leftChildren, node.leaf, node.parent}
	right := &treeNode{rightKeys, rightChildren, node.leaf, node.parent}
	parent := node.parent
	if parent == nil {
		// new parent
		parent = &treeNode{[]int{upKey}, []*treeNode{left, right}, false, nil}
		return parent
	}
	// up node
	idx, ln := 0, len(parent.keys)
	for idx < ln {
		if parent.keys[idx] > upKey {
			break
		}
	}
	parent.keys = append(parent.keys, nilKey)
	parent.children = append(parent.children, nil)
	idx2 := ln
	for idx2 > idx {
		parent.keys[idx2] = parent.keys[idx2-1]
		parent.children[idx2+1] = parent.children[idx2]
	}
	parent.keys[idx2], parent.children[idx2], parent.children[idx+1] = upKey, left, right
	return split(parent, t)
}
