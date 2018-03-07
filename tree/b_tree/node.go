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
