package b_tree

type treeNode struct {
	keys     []int
	children []*treeNode
	leaf     bool
	parent   *treeNode
}

const nilKey = 0

func findInKeys(keys []int, key int) (int, bool) {
	idx, ln := 0, len(keys)
	for idx < ln {
		if keys[idx] == key {
			return idx, true
		} else if keys[idx] > key {
			break
		}
	}
	return idx, false
}

// // just insert the key and nothing will happen to children
func insertKeyIntoLeafNode(node *treeNode, key int) int {
	idx, ln := 0, len(node.keys)
	node.keys = append(node.keys, nilKey)
	idx2 := ln
	for idx2 > idx {
		node.keys[idx2] = node.keys[idx2-1]
		idx2 -= 1
	}
	node.keys[idx] = key
	return idx
}

func indexOfKey(node *treeNode, key int) int {
	for i := 0; i < len(node.keys); i++ {
		if node.keys[i] == key {
			return i
		}
	}
	return -1
}

func indexOfChild(parent, child *treeNode) int {
	for i := 0; i < len(parent.children); i++ {
		if parent.children[i] == child {
			return i
		}
	}
	return -1
}

// just delete the key and nothing will happen to children
func deleteKeyWithIndex(node *treeNode, index int) {
	index += 1
	ln := len(node.keys)
	for index < ln {
		node.keys[index-1] = node.keys[index]
		index += 1
	}
	node.keys = node.keys[0 : ln-1]
}

// just delete the key and nothing will happen to children
func deleteKey(node *treeNode, key int) {
	idx := indexOfKey(node, key)
	if idx < 0 {
		return
	}
	deleteKeyWithIndex(node, idx)
}

// combine two children whose indices are 'index' and 'index+1'
func combine(node *treeNode, index, t int) {
	// combine
	key := node.keys[index]
	deleteKeyWithIndex(node, index)
	left, right := node.children[index], node.children[index+1]
	leftLn, rightLn := len(left.keys), len(right.keys)
	newKeys := make([]int, leftLn+rightLn+1)
	for i := 0; i < leftLn; i++ {
		newKeys[i] = left.keys[i]
	}
	newKeys[leftLn] = key
	for i := 0; i < rightLn; i++ {
		newKeys[i+leftLn+1] = right.keys[i]
	}
	newChildren := make([]*treeNode, leftLn+1+rightLn+1)
	for i := 0; i <= leftLn; i++ {
		newChildren[i] = left.children[i]
	}
	for i := 0; i <= rightLn; i++ {
		newChildren[i+leftLn+1] = right.children[i]
	}
	newNode := &treeNode{newKeys, newChildren, left.leaf, node}
	node.children[index] = newNode
	ln := len(node.children)
	for i := index + 1; i < ln-1; i++ {
		node.children[i] = node.children[i+1]
	}
	node.children = node.children[0 : ln-1]
	// node's length of keys may be less than t-1
	if node.parent != nil && len(node.keys) <= t-1 {
		pIdx := indexOfChild(node.parent, node)
		if pIdx > 0 {
			left := node.parent.children[pIdx-1]
			if len(left.keys) > t-1 {
				// borrow
				borrowFromLeft(node.parent, pIdx)
			} else {
				// combine
				combine(node.parent, pIdx-1, t)
			}
		} else {
			// the number of children of a node must be at least 2
			right := node.parent.children[pIdx+1]
			if len(right.keys) > t-1 {
				// borrow
				borrowFromRight(node.parent, pIdx)
			} else {
				// combine
				combine(node.parent, pIdx, t)
			}
		}
	}

}

// index: the index of node in parent's children
func borrowFromLeft(parent *treeNode, index int) {
	left, node := parent.children[index-1], parent.children[index]
	node.keys = append([]int{parent.keys[index-1]}, node.keys...)
	leftLn := len(left.keys)
	node.children = append([]*treeNode{left.children[leftLn]}, node.children...)
	parent.keys[index-1] = left.keys[leftLn-1]
	left.keys = left.keys[0 : leftLn-1]
	left.children = left.children[0:leftLn]
}

// index: the index of node in parent's children
func borrowFromRight(parent *treeNode, index int) {
	right, node := parent.children[index+1], parent.children[index]
	node.keys = append(node.keys, parent.keys[index])
	node.children = append(node.children, right.children[0])
	parent.keys[index] = right.keys[0]
	right.keys = right.keys[1:]
	right.children = right.children[1:]
}
