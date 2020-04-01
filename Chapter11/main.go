package main

import (
	"fmt"
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// samplePushPopHeap()
	sortSortedArrays()
}

func sortSortedArrays() {
	arrays := [][]int{
		[]int{0, 6},
		[]int{0, 6, 28},
		[]int{3, 5, 8},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	fmt.Println(arrays)
	fmt.Println(mergeSortedArrays(arrays))
}

// func samplePushPopHeap() {
// 	h := &IntHeap{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
// }
