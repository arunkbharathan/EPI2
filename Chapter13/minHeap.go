// This example demonstrates an integer heap built using the heap interface.
package main

type stringCounter struct {
	Str      *string
	StrCount int
}

// An stringCounterHeap is a min-heap of string counts.
type stringCounterHeap []*stringCounter

func (h stringCounterHeap) Len() int {
	v := len(h)
	return v
}
func (h stringCounterHeap) Less(i, j int) bool {
	a := h[i].StrCount
	b := h[j].StrCount
	return a < b
}
func (h stringCounterHeap) Swap(i, j int) {
	a, b := h[j], h[i]
	h[i], h[j] = a, b
}

// Push push to heap
func (h *stringCounterHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*stringCounter))
}

//Pop pop from heap
func (h *stringCounterHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
