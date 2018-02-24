package BH

type BinomialHeap struct {
	Head *BinomialTree
	Next *BinomialHeap
}

// MakeBinomialHeap make a new binomial heap
func MakeBinomialHeap() *BinomialHeap {
	res := new(BinomialHeap)
	return res
}

// TODO Modify all below
// BinomialHeapMinimum find the minimum key
func BinomialHeapMinimum(h *BinomialTree) *BinomialTree {
	var (
		x   *BinomialTree = h
		y   *BinomialTree = nil
		min               = 1<<32 - 1
	)
	for x != nil {
		if x.key < min {
			min = x.key
			y = x
		}
		x = x.Sibling
	}
	return y
}

// BinomialHeapUnion union 2 binomial heaps
func BinomialHeapUnion(h1, h2 *BinomialTree) *BinomialTree {
	h := MakeBinomialHeap()
	h.Head = binomialHeapMerge(h1, h2)
	if h.Head == nil {
		return h
	}

	var (
		prev *BinomialTree = nil
		x    *BinomialTree = h.Head
		next *BinomialTree = x.Sibling
	)
	for next != nil {
		if x.Degree != next.Degree || next.Sibling != nil && next.Sibling.Degree == x.Degree {
			// case 1 and 2
			prev = x
			x = next
		} else if x.key <= next.key {
			// case 3
			x.Sibling = next.Sibling
			binomialLink(next, x)
		} else if prev == nil {
			// case 4
			h.Head = next
		} else {
			// case 4
			prev.Sibling = next
			binomialLink(x, next)
			x = next
		}
		next = x.Sibling
	}
	return h
}

// binomialLink y to z
func binomialLink(y, z *BinomialTree) {
	y.Parent = z
	y.Sibling = z.Child
	z.Child = y
	z.Degree += 1
}

// binomialHeapMerge merge 2 binomial heaps
func binomialHeapMerge(h1, h2 *BinomialTree) *BinomialTree {

	var (
		c1 *BinomialTree = h1
		c2 *BinomialTree = h2
	)

}
