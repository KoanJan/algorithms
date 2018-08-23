package base

import "testing"

type testPItem int

func (t testPItem) Priority() int {
	return int(t)
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	pq.Enqueue(testPItem(1))
	pq.Enqueue(testPItem(3))
	pq.Enqueue(testPItem(2))
	for !pq.IsEmpty() {
		t.Log(pq.Dequeue())
	}
}
