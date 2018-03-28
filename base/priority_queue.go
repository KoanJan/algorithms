package base

import "container/heap"

type PriorityQueue struct {
	data *priorityQueue
}

func NewPriorityQueue() *PriorityQueue {
	queue := new(PriorityQueue)
	queue.data = new(priorityQueue)
	*queue.data = make([]PriorityItem, 0)
	return queue
}

func (queue *PriorityQueue) Enqueue(item interface{}) {
	heap.Push(queue.data, item)
}

func (queue *PriorityQueue) Dequeue() interface{} {
	return heap.Pop(queue.data).(PriorityItem)
}

func (queue *PriorityQueue) IsEmpty() bool {
	return queue.data.Len() == 0
}

type PriorityItem interface {
	Priority() int
}

type priorityQueue []PriorityItem

func (queue priorityQueue) Len() int {
	return len(queue)
}

func (queue priorityQueue) Less(i, j int) bool {
	return queue[i].Priority() < queue[j].Priority()
}

func (queue priorityQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}

func (queue *priorityQueue) Push(x interface{}) {
	*queue = append(*queue, x.(PriorityItem))
}

func (queue *priorityQueue) Pop() interface{} {
	n := len(*queue)
	x := (*queue)[n-1]
	*queue = (*queue)[0 : n-1]
	return x
}
