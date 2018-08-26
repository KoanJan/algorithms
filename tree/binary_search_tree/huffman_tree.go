package binary_search_tree

import (
	"algorithms/heap"
)

type HuffmanTree struct {
	key         string
	w           int
	left, right *HuffmanTree
}

type HuffmanNode struct {
	Key string
	W   int
}

func (tree *HuffmanTree) Value() int {
	return -tree.w
}

func NewHuffmanTree(nodes ...HuffmanNode) *HuffmanTree {
	if len(nodes) == 0 {
		return nil
	}
	// Sorting by a heap is simple(each operation costs O(logn)).
	// Suggestion:
	// To increase the efficiency, you can keep a sorted
	// node linked list, and find the right location by
	// binary search while insertion, which cost O(logn).
	// And each popping costs O(1).
	h := heap.NewBinaryHeap()
	for _, n := range nodes {
		h.Insert(&HuffmanTree{key: n.Key, w: n.W})
	}
	for len(*h) > 1 {
		t1 := h.Pop().(*HuffmanTree)
		t2 := h.Pop().(*HuffmanTree)
		h.Insert(&HuffmanTree{w: t1.w + t2.w})
	}
	return h.Pop().(*HuffmanTree)
}
