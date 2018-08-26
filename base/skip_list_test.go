package base

import "testing"

type testV int

func (v testV) Value() int {
	return int(v)
}

func TestSkipList(t *testing.T) {
	list := NewSkipList(5)
	// add
	list.Add(testV(6))
	list.Add(testV(6))
	list.Add(testV(5))
	list.Add(testV(2))
	list.Add(testV(1))
	list.Add(testV(4))
	list.Add(testV(7))
	list.Add(testV(7))
	list.Add(testV(4))
	t.Log(list)
	// find
	t.Log(list.Find(testV(6)))
	t.Log(list.Find(testV(2)))
	// delete
	list.Delete(testV(6))
	t.Log(list)
	t.Log(list.Find(testV(6)))
	t.Log(list)
}
