package hashmap

const (
	initialCapacity = 1 << 8
	maximumCapacity = 1 << 30
	loadFactor      = 0.75
)

// HashMap hash map
type hashMap struct {
	capacity  int // size of HashMap
	threshold int
	buckets   []*linkedList // store linked list
}

func (h *hashMap) Len() int {
	return h.capacity
}

func (h *hashMap) Cap() int {
	return h.threshold
}

// Put val into map index by key
func (h *hashMap) Put(k string, v interface{}) {

	if h.needExpansion() {
		h.expand()
	}

	index := hashCode(k) % h.capacity
	if h.buckets[index] == nil {
		// init linked list
		h.buckets[index] = &linkedList{Key: &k, Val: v}
	} else {
		l := h.buckets[index]
		for *l.Key != k {
			if l.Next == nil {
				break
			}
			l = l.Next
		}
		if *l.Key == k {
			// update
			l.Val = v
		} else {
			// append
			l.Next = &linkedList{Key: &k, Val: v}
		}
	}
	return
}

// Get val by key
func (h *hashMap) Get(k string) (interface{}, bool) {
	index := hashCode(k) % h.capacity
	if h.buckets[index] == nil {
		return nil, false
	}
	l := h.buckets[index]
	for *l.Key != k {
		if l = l.Next; l == nil {
			return nil, false
		}
	}
	return l.Val, true
}

func (h *hashMap) needExpansion() bool {
	return h.threshold < calcThreshold(h.capacity, loadFactor)
}

func (h *hashMap) expand() {
	// check if the capacity permitted
	if h.capacity > maximumCapacity>>1 {
		panic("too many values for hashmap")
	}
	// pre-expand
	size := h.capacity << 1
	newBuckets := make([]*linkedList, size)
	// copy data
	for _, bucket := range h.buckets {
		for bucket != nil {
			index := hashCode(*bucket.Key) % size
			// insert to the head of linked list
			newBuckets[index] = &linkedList{bucket.Key, bucket.Val, newBuckets[index]}
		}
	}
	// update info
	h.threshold = calcThreshold(size, loadFactor)
	h.buckets = newBuckets
}

func NewHashMap() Map {
	return &hashMap{initialCapacity, calcThreshold(initialCapacity, loadFactor), make([]*linkedList, initialCapacity)}
}

func calcThreshold(size int, lf float64) int {
	return int(float64(size) * lf)
}
