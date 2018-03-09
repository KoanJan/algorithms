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
	if tree.node == nil {
		tree.node = &treeNode{[]int{key}, nil, true, nil}
		return
	}
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
		leftChildren, rightChildren = node.children[0:t+1], node.children[t+1:]
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
		parent.keys[idx2], parent.children[idx2+1] = parent.keys[idx2-1], parent.children[idx2]
		idx2 -= 1
	}
	parent.keys[idx2], parent.children[idx2], parent.children[idx+1] = upKey, left, right
	return split(parent, t)
}

func (tree *BTree) Delete(key int) {
	node, exist := findInBTreeNode(tree.node, key)
	if exist {
		node = deleteKeyFromNode(node, key, tree.t)
		for node.parent != nil {
			node = node.parent
		}
		tree.node = node
	}
}

func deleteKeyFromNode(node *treeNode, key, t int) *treeNode {
	//
	idx := indexOfKey(node, key)
	if node.parent == nil || len(node.keys) > t-1 {
		if node.leaf {
			deleteKey(node, key)
			return node
		} else {
			if len(node.children[idx].keys) > t-1 {
				// swap and recursively delete
				node.keys[idx] = node.children[idx].keys[len(node.children[idx].keys)-1]
				return deleteKeyFromNode(node.children[idx], node.keys[idx], t)
			} else if len(node.children[idx+1].keys) > t-1 {
				// swap and recursively delete
				node.keys[idx] = node.children[idx+1].keys[0]
				return deleteKeyFromNode(node.children[idx+1], node.keys[idx], t)
			} else {
				// combine
				combine(node, idx, t)
				return deleteKeyFromNode(node.children[idx], key, t)
			}
		}
	} else {
		// balance
		pIdx := indexOfChild(node.parent, node)
		if pIdx > 0 && len(node.parent.children[pIdx-1].keys) > t-1 {
			// borrow from left sibling
			borrowFromLeft(node.parent, pIdx)
		} else if pIdx < len(node.parent.children)-1 && len(node.parent.children[pIdx+1].keys) > t-1 {
			// borrow from right sibling
			borrowFromRight(node.parent, pIdx)
		} else if pIdx > 0 {
			// combine with left sibling
			combine(node.parent, pIdx-1, t)
			node = node.parent.children[pIdx-1]
		} else {
			// combine with right sibling
			combine(node.parent, pIdx, t)
			node = node.parent.children[pIdx]
		}
		return deleteKeyFromNode(node, key, t)
	}
}
