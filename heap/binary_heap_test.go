package heap

import "testing"

type v int

func (v v) Value() int {
	return int(v)
}

func TestBinaryHeap(t *testing.T) {
	h := NewBinaryHeap()
	h.Insert(v(1))
	h.Insert(v(2))
	h.Insert(v(3))
	h.Insert(v(4))
	h.Insert(v(5))
	h.Insert(v(6))
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
	t.Log(h.Pop())
	t.Log(*h)
}
