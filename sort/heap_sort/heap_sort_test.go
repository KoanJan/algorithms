package heap_sort

import "testing"

func TestHeapSort(t *testing.T) {
	a := []int{9, 7, 4, 6, 3, 1, 5, 2, 8}
	HeapSort(a)
	t.Log(a)
}

func BenchmarkHeapSort(b *testing.B) {
	a := []int{9, 7, 4, 6, 3, 1, 5, 2, 8}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HeapSort(a)
	}
}
