package base

// queue
type Queue interface {

	// enqueue
	Enqueue(interface{})

	// dequeue
	Dequeue() interface{}

	// is empty
	IsEmpty() bool
}
