package main

// BNode B-Tree node
type BNode struct {
	IsLeaf bool

	children     []*BNode // length is n+1
	keys         []int    // length is n
	maxChildrens int      // maximum of the number of children
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
	if index < 0 || index > len(b.keys-1) {
		panic("Out of index")
	}
	return b.keys[index]
}

// Key get the index of k in keys
func (b *BNode) Key(k int) int {
	var (
		l   int = 0
		r   int = len(b.keys)
		mid int = len(b.keys) / 2
	)
	for l < r {
		if mid == l {
			l = r
		}
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
	}
	return -1
}

// Search
func (b *BNode) Search(k int) (*BNode, int) {
	if b.Key(k) != -1 {
		return b, b.Key(k)
	}
	if b.IsLeaf {
		return nil, -1
	}
	var (
		r int = 0
	)
	for index, key := range b.keys {
		if key > k {
			r = index
			break
		}
	}
	return b.children[r].Search(k)
}
