package main

type StackQueue struct {
	store, cache *Stack
}

func (this *StackQueue) Enqueue(v interface{}) {
	// lazy
	for !this.cache.IsEmpty() {
		this.store.Push(this.cache.Pop())
	}
	this.store.Push(v)
}

func (this *StackQueue) Dequeue() interface{} {
	// lazy
	for !this.store.IsEmpty() {
		this.cache.Push(this.store.Pop())
	}
	return this.cache.Pop()
}

func (this *StackQueue) IsEmpty() bool {
	return this.store.IsEmpty()
}

func NewStackQueue() *StackQueue {
	return &StackQueue{store: NewStack(), cache: NewStack()}
}
