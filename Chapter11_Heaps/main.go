package main

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	// samplePushPopHeap()
	// sortSortedArrays()
	// sortAlmostSortedArray()
	// findKclosestStars()
	findRunningMedian()
}

func findRunningMedian() {
	streamingMedian()
}

// func findKclosestStars() {
// 	fmt.Println(findNclosestStars(5))
// }

// func sortAlmostSortedArray() {
// 	array := []int{3, -1, 2, 6, 4, 5, 8, 6, 7, 12, 9, 10, 14, 12, 13, 20, 15}
// 	fmt.Println(array)
// 	fmt.Println(sortArray(array, 4))
// }

// func sortSortedArrays() {
// 	arrays := [][]int{
// 		[]int{0, 6},
// 		[]int{0, 6, 28},
// 		[]int{3, 5, 7},
// 		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 	}
// 	fmt.Println(arrays)
// 	fmt.Println(mergeSortedArrays(arrays))
// }

// func samplePushPopHeap() {
// 	h := &IntHeap{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
// }
