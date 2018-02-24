package main

// 虚拟节点
type virtualNode struct {
	val string
}

// 真实节点
type node struct {
	val string
}

// 哈希环
type HashRing struct {
	virtualMap map[virtualNode]node
}

// 添加节点
func (h *HashRing) AddNode(val string) {

}
