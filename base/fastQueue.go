package main

type FastQueue struct {
	head, tail *LinkedList
}

func (this *FastQueue) Enqueue(v interface{}) {
	this.tail.Next = &LinkedList{V: v}
	this.tail = this.tail.Next
}

// unsafe
func (this *FastQueue) Dequeue() interface{} {
	v := this.head.Next.V
	this.head.Next = this.head.Next.Next
	return v
}

func (this *FastQueue) IsEmpty() bool {
	return this.head.Next == nil
}

func NewFastQueue() *FastQueue {
	queue := new(FastQueue)
	queue.head = new(LinkedList)
	queue.tail = queue.head
	return queue
}
