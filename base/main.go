package main

import (
	"fmt"
)

func main() {
	// queue
	queue := NewFastQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("queue dequeue: ")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	// stack
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("stack pop: ")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
