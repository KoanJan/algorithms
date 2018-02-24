package main

import (
	"fmt"
)

const T int = 2

// BNode B-Tree node
type BNode struct {
	IsLeaf bool

	children []*BNode // length is n+1
	keys     []int    // length is n
}

// CountKeys
func (b *BNode) CountKeys() int {
	return len(b.keys)
}

// H
func (b *BNode) H() int {
	// TODO: return height of B-tree
	return 0
}

// Keys get key by index
func (b *BNode) Keys(index int) int {
	if index < 0 || index > len(b.keys)-1 {
		panic("Out of index")
	}
	return b.keys[index]
}

// Key get the index of k in keys
func (b *BNode) Key(k int) int {
	var (
		l   int = 0
		r   int = len(b.keys) - 1
		mid int = (len(b.keys) - 1) / 2
	)
	for l < r {
		m := b.keys[mid]
		if m == k {
			return mid
		}
		if m < k {
			l = mid
		} else {
			r = mid
		}
		mid = (l + r) / 2
		if mid == l {
			if b.keys[r] == k {
				return r
			}
			break
		}
	}
	return -1
}

// ChildIndex
func (b *BNode) ChildIndex(k int) int {
	// in b
	if b.Key(k) != -1 {
		return -1
	}
	var i int = 0
	for idx, v := range b.keys {
		if v > k {
			i = idx
			break
		}
	}
	return i
}

// Search
func (b *BNode) Search(k int) (*BNode, int) {
	if b.Key(k) != -1 {
		return b, b.Key(k)
	}
	if b.IsLeaf {
		return nil, -1
	}

	return b.children[b.ChildIndex(k)].Search(k)
}

// Insert
func (b *BNode) Insert(k int) {

	var (
		c      *BNode = b
		parent *BNode = nil
	)

	// locate
	if c.IsLeaf {
		// h = 1
		if len(c.children) == 2*T-1 {
			// split
			splitBNode(c, parent)
			parent = c
			c = c.children[c.ChildIndex(k)]
		}
	} else {
		// h > 1
		if len(c.children) == 2*T-1 {
			// split
			splitBNode(c, parent)
			parent = c
			c = c.children[c.ChildIndex(k)]
		}
		for (!c.IsLeaf) && len(c.children) == 2*T-1 {
			// k existed
			if c.Key(k) != -1 {
				return
			}
			parent = c
			c = c.children[c.ChildIndex(k)]
			if len(c.children) == 2*T-1 {
				splitBNode(c, parent)
			}
		}
	}

	// k existed
	if c.Key(k) != -1 {
		return
	}

	// insert
	var (
		l   int = 0
		r   int = c.CountKeys()
		mid int = (c.CountKeys() - 1) / 2
	)
	if r < 0 {
		r = 0
	}
	for l < r {
		if c.Keys(mid) > k {
			r = mid
		} else {
			l = mid
		}
		mid = (l + r) / 2
		if mid == l {
			l = r
			break
		}
	}
	newKeys := []int{}
	for i := 0; i < r; i++ {
		newKeys = append(newKeys, c.keys[i])
	}
	newKeys = append(newKeys, k)
	for i := r; i < c.CountKeys(); i++ {
		newKeys = append(newKeys, c.keys[i])
	}
	c.keys = newKeys
	c.children = append(c.children, nil)
}

func splitBNode(node *BNode, parent ...*BNode) {

	var p *BNode
	if len(parent) > 0 {
		p = parent[0]
	}

	index := node.CountKeys() / 2
	parentKey := node.Keys(index)

	// left child
	left := &BNode{}
	for i := 0; i < index; i++ {
		left.keys = append(left.keys, node.keys[i])
		left.children = append(left.children, node.children[i])
	}
	left.children = append(left.children, node.children[index])
	left.IsLeaf = true
	for _, child := range left.children {
		if child != nil {
			left.IsLeaf = false
			break
		}
	}

	// right child
	right := &BNode{}
	for i := index + 1; i < node.CountKeys(); i++ {
		right.keys = append(right.keys, node.keys[i])
		right.children = append(right.children, node.children[i])
	}
	right.children = append(right.children, node.children[node.CountKeys()])
	right.IsLeaf = true
	for _, child := range right.children {
		if child != nil {
			right.IsLeaf = false
			break
		}
	}

	// if root
	if p == nil {
		node.keys = []int{parentKey}
		node.children = []*BNode{left, right}
		node.IsLeaf = false
	} else {
		l := 0
		r := p.CountKeys() - 1
		mid := (p.CountKeys() - 1) / 2
		for l < r {
			// parentKey won't exsited in p
			if p.Keys(mid) < parentKey {
				l = mid
			} else {
				r = mid
			}
			mid = (l + r) / 2
			// reach where to insert
			if r-l == 1 {
				newParentKeys := []int{}
				newParentChildren := []*BNode{}
				for i := 0; i < r; i++ {
					newParentKeys = append(newParentKeys, p.Keys(i))
					newParentChildren = append(newParentChildren, p.children[i])
				}
				newParentKeys = append(newParentKeys, parentKey)
				newParentChildren = append(newParentChildren, left)
				newParentChildren = append(newParentChildren, right)
				for i := r; i < p.CountKeys(); i++ {
					newParentKeys = append(newParentKeys, p.Keys(i))
					newParentChildren = append(newParentChildren, p.children[i+1])
				}
				p.keys = newParentKeys
				p.children = newParentChildren
				p.IsLeaf = false
				break
			}
		}
	}
}

// NewBNode new B-Node
func NewBNode(keys ...int) *BNode {
	node := &BNode{IsLeaf: true}
	if len(keys) != 0 {
		if len(keys) >= 2*T-1 {
			panic("keys is too many")
		}

		// simple O(n^2) sort
		for i := 0; i < len(keys); i++ {
			for j := len(keys) - 1; j > i; j-- {
				if keys[i] > keys[j] {
					tmp := keys[i]
					keys[i] = keys[j]
					keys[j] = tmp
				}
			}
		}

		// add keys and init children
		node.children = []*BNode{nil}
		for _, key := range keys {
			node.keys = append(node.keys, key)
			node.children = append(node.children, nil)
		}
	}
	return node
}

// String format output
func (b *BNode) String() string {
	return fmt.Sprintf("{keys: %v, children: %v, isLeaf: %v}", b.keys, b.children, b.IsLeaf)
}

// Delete delete key
func (b *BNode) Delete(k int) {
	// TODO
	node := b.Search(k)
	if node.IsLeaf {
		//  simple linear find and delete
		newKeys := []int{}
		for i := 0; i < node.CountKeys(); i++ {
			if node.keys[i] != k {
				newKeys = append(newKeys, node.keys[i])
			}
		}
		node.keys = newKeys
		node.children = node.children[1:]
	} else {
		//
	}
}
