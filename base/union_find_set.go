package base

const nilIndex = -1

type UnionFindSet struct {
	index map[interface{}]int
	data  []interface{}
	ref   []int
}

func NewUnionFindSet(data ...interface{}) *UnionFindSet {
	set := new(UnionFindSet)
	ln := len(data)
	set.index = make(map[interface{}]int)
	set.data = make([]interface{}, ln)
	set.ref = make([]int, ln)
	if ln > 0 {
		for i := 0; i < ln; i++ {
			set.index[data[i]] = i
			set.data[i] = data[i]
			set.ref[i] = nilIndex
		}
	}
	return set
}

func (set *UnionFindSet) SetParent(child, parent interface{}) {
	set.ref[set.index[child]] = set.index[parent]
}

func (set *UnionFindSet) Find(node interface{}) interface{} {
	return set.data[set.find(set.index[node])]
}

func (set *UnionFindSet) find(index int) int {
	parent := set.ref[index]
	if parent == nilIndex {
		return index
	}
	return set.find(parent)
}
