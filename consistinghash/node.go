package consistinghash

import "strconv"

const vNodeSize = 16

type Node struct {
	ip   string
	name string
	sign uint32

	data map[string]interface{}

	vNodes []*vNode
}

type vNode struct {
	sign     uint32
	realNode *Node
}

func (node *Node) Sign() uint32 {
	return node.sign
}

func (node *Node) Set(k string, v interface{}) {
	node.data[k] = v
}

func (node *Node) Get(k string) (interface{}, bool) {
	v, existed := node.data[k]
	return v, existed
}

func (node *Node) Del(k string) {
	delete(node.data, k)
}

func NewNode(ip, name string) *Node {
	node := &Node{ip: ip, name: name}
	node.sign = Hash(node.ip + node.name)
	node.data = make(map[string]interface{})
	node.vNodes = make([]*vNode, vNodeSize)
	for i := 0; i < vNodeSize; i++ {
		node.vNodes[i] = &vNode{Hash(node.ip + node.name + strconv.Itoa(i)), node}
	}
	return node
}
