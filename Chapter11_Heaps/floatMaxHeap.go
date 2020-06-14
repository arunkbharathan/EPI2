// This example demonstrates an float heap built using the heap interface.
package main

// An FloatMaxHeap is a max-heap of floats.
type FloatMaxHeap []float64

func (h FloatMaxHeap) Len() int           { return len(h) }
func (h FloatMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h FloatMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push push to heap
func (h *FloatMaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float64))
}

//Pop pop from heap
func (h *FloatMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
