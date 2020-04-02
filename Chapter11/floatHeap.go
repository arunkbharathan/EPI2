// This example demonstrates an float heap built using the heap interface.
package main

// An FloatHeap is a min-heap of floats.
type FloatHeap []float64

func (h FloatHeap) Len() int           { return len(h) }
func (h FloatHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h FloatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push push to heap
func (h *FloatHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float64))
}

//Pop pop from heap
func (h *FloatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
