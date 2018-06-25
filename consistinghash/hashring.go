package consistinghash

import "sort"

type HashRing []*vNode

func (ring HashRing) Len() int {
	return len(ring)
}

func (ring HashRing) Less(i, j int) bool {
	return ring[i].sign < ring[j].sign
}

func (ring HashRing) Swap(i, j int) {
	ring[i], ring[j] = ring[j], ring[i]
}

func NewHashRing() *HashRing {
	ring := &HashRing{}
	return ring
}

func (ring *HashRing) AddNode(node *Node) {
	*ring = append(*ring, node.vNodes...)
	sort.Sort(*ring)
}

func (ring *HashRing) DelNode(node *Node) {

	toDel := make(map[*vNode]bool)
	for _, vNode := range node.vNodes {
		toDel[vNode] = true
	}
	newSize := ring.Len() - len(node.vNodes)
	if newSize < 0 {
		newSize = 0
	}
	newRing := make([]*vNode, newSize)
	for _, oldVNode := range *ring {
		if _, existed := toDel[oldVNode]; !existed {
			newRing = append(newRing, oldVNode)
		}
	}
	*ring = newRing
	sort.Sort(*ring)
}

func (ring *HashRing) Set(k string, v interface{}) {
	findFixedNode(ring, k).realNode.Set(k, v)
}

func (ring *HashRing) Get(k string) (interface{}, bool) {
	return findFixedNode(ring, k).realNode.Get(k)
}

func (ring *HashRing) Del(k string) {
	findFixedNode(ring, k).realNode.Del(k)
}

func findFixedNode(ring *HashRing, key string) *vNode {
	hashVal := Hash(key)
	i := 0
	for i < ring.Len() {
		if hashVal <= (*ring)[i].sign {
			break
		}
		i++
	}
	if i == ring.Len() {
		i = 0
	}
	return (*ring)[i]
}
