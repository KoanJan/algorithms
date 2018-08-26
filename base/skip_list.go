package base

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// SkipList is an optimism data structure from linked list.
// It has ascendant search than linked list.
type SkipList struct {
	head  *skipNode
	maxLv int // maximum of level count
	lv    int // current level count
}

type skipNode struct {
	v Valuable
	// The 'next' field contains all next node pointers of each level
	// this node contains, so that duplicate 'v' is avoided.
	// eg. next[1] means the next node of this with level 2(1+1).
	next []*skipNode
}

func NewSkipList(maxLv int) *SkipList {
	list := &SkipList{maxLv: maxLv}
	return list
}

func (list *SkipList) String() string {
	if list.head == nil {
		return ""
	}
	buf := bytes.NewBufferString("")
	for i := list.lv - 1; i >= 0; i-- {
		for cur := list.head; cur != nil; cur = cur.next[i] {
			buf.WriteString(fmt.Sprintf("%d ", cur.v.Value()))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (list *SkipList) Add(v Valuable) {
	cur := list.head
	if cur == nil {
		// empty skip list
		list.head = &skipNode{v: v, next: make([]*skipNode, 1)}
		list.lv = 1
		return
	}
	val := v.Value()
	headVal := cur.v.Value()
	if headVal > val {
		// insert before head
		k := randomLv(list.maxLv)
		if k < list.lv {
			k = list.lv
		}
		newHead := &skipNode{v: v, next: make([]*skipNode, k)}
		for i := 0; i < list.lv; i++ {
			newHead.next[i] = list.head
		}
		list.head = newHead
	} else if headVal < val {
		// find existed node or pre nodes
		pre := make([]*skipNode, list.lv)
		for i := list.lv - 1; i >= 0; i-- {
			for cur.next != nil && cur.next[i] != nil {
				nextV := cur.next[i].v.Value()
				if nextV == val {
					// exist
					return
				}
				if nextV > val {
					break
				}
				cur = cur.next[i]
			}
			pre[i] = cur
		}
		// insert new node
		k := randomLv(list.maxLv)
		newNode := &skipNode{v: v, next: make([]*skipNode, k)}
		for i := 0; i < list.lv; i++ {
			newNode.next[i] = pre[i].next[i]
			pre[i].next[i] = newNode
		}
		// check if new lv is need
		if k > list.lv {
			for i := list.lv; i < k; i++ {
				list.head.next = append(list.head.next, newNode)
			}
			list.lv = k
		}
	}
	// headVal == val
}

func (list *SkipList) Delete(v Valuable) {
	cur := list.head
	// empty skip list
	if cur == nil {
		return
	}
	val := v.Value()
	if cur.v.Value() > val {
		// not exist
		return
	}
	if cur.v.Value() == val {
		// delete head
		// fix lv of the new head
		for i := len(cur.next[0].next); i < list.lv; i++ {
			cur.next[0].next = append(cur.next[0].next, cur.next[i])
		}
		list.head = cur.next[0]
		return
	}
	for i := list.lv - 1; i >= 0; i-- {
		for cur.next != nil && cur.next[i] != nil {
			nextV := cur.next[i].v.Value()
			if nextV == val {
				// delete
				// found the node to delete
				toDelete := cur.next[i]
				// delete all links
				for j := i; j >= 0; j-- {
					for cur.next[j] != toDelete {
						cur = cur.next[j]
					}
					cur.next[j] = toDelete.next[j]
				}
				return
			}
			if nextV > val {
				break
			}
			cur = cur.next[i]
		}
	}
	// not exist
}

func (list *SkipList) Find(v Valuable) Valuable {
	cur := list.head
	// empty skip list
	if cur == nil {
		return nil
	}
	val := v.Value()
	if cur.v.Value() >= val {
		return cur.v
	}
	for i := list.lv - 1; i >= 0; i-- {
		for cur.next != nil && cur.next[i] != nil {
			nextV := cur.next[i].v.Value()
			if nextV == val {
				return cur.next[i].v
			}
			if nextV > val {
				break
			}
			cur = cur.next[i]
		}
	}
	return nil
}

// A random lv not larger than max will be generated.
func randomLv(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	k := 1
	for k < max && r.Intn(2) == 1 {
		k++
	}
	return k
}
