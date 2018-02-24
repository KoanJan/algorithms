package main

type Stack struct {
	data []interface{}
}

func (this *Stack) Push(v interface{}) {
	this.data = append(this.data, v)
}

func (this *Stack) Pop() interface{} {
	oldLength := len(this.data)
	v := this.data[oldLength-1]
	this.data = this.data[0 : oldLength-1]
	return v
}

func (this *Stack) IsEmpty() bool {
	return len(this.data) == 0
}

func NewStack() *Stack {
	return &Stack{data: []interface{}{}}
}
