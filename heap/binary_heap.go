package heap

import "algorithms/base"

type BinaryHeap []base.Valuable

func (h *BinaryHeap) Insert(v base.Valuable) {
	idx := len(*h)
	var p int
	*h = append(*h, v)
	for idx > 0 {
		p = (idx - 1) >> 1
		if v.Value() > (*h)[p].Value() {
			(*h)[idx], (*h)[p] = (*h)[p], (*h)[idx]
			idx = p
			continue
		}
		break
	}
}

func (h *BinaryHeap) Pop() base.Valuable {
	n := len(*h)
	if n == 0 {
		return nil
	}
	v := (*h)[0]
	newN := n - 1
	(*h)[0] = (*h)[newN]
	*h = (*h)[:newN]
	idx := 0
	var lCIdx = 1
	for lCIdx < newN {
		if lCIdx+1 < newN && (*h)[lCIdx].Value() < (*h)[lCIdx+1].Value() {
			if (*h)[idx].Value() < (*h)[lCIdx+1].Value() {
				(*h)[idx], (*h)[lCIdx+1] = (*h)[lCIdx+1], (*h)[idx]
				idx = lCIdx + 1
				lCIdx = (idx << 1) | 1
				continue
			}
		} else if (*h)[idx].Value() < (*h)[lCIdx].Value() {
			(*h)[idx], (*h)[lCIdx] = (*h)[lCIdx], (*h)[idx]
			idx = lCIdx
			lCIdx = (idx << 1) | 1
			continue
		}
		break
	}
	return v
}

func NewBinaryHeap() *BinaryHeap {
	h := make(BinaryHeap, 0, 1)
	return &h
}
