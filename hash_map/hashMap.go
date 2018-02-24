package main

type Map interface {
	Put(string, interface{})
	Get(string) (interface{}, bool)
	// Len() int
	// Cap() int
}

// HashMap hash map
type hashMap struct {
	m    int           // size of HashMap
	vals []*linkedList // store linked list
}

// Put val into map index by key
func (h *hashMap) Put(k string, v interface{}) {
	index := hashCode(k) % h.m
	if h.vals[index] == nil {
		// init linked list
		h.vals[index] = &linkedList{Key: &k, Val: v}
	} else {
		l := h.vals[index]
		for *l.Key != k {
			if l.Next == nil {
				break
			}
			l = l.Next
		}
		if *l.Key == k {
			// update
			l.Val = v
		} else {
			// append
			l.Next = &linkedList{Key: &k, Val: v}
		}
	}
	return
}

// Get val by key
func (h *hashMap) Get(k string) (interface{}, bool) {
	index := hashCode(k) % h.m
	if h.vals[index] == nil {
		return nil, false
	}
	l := h.vals[index]
	for *l.Key != k {
		if l = l.Next; l == nil {
			return nil, false
		}
	}
	return l.Val, true
}

func NewHashMap(m ...int) Map {
	_m := 1 << 16
	if len(m) > 0 {
		if m[0] < _m && m[0] > 0 {
			_m = m[0]
		}
	}
	vals := make([]*linkedList, _m)
	return &hashMap{m: _m, vals: vals}
}
