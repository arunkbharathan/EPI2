package main

import "container/heap"

func sortArray(arr []int, chunk int) []int {
	result := []int{}
	h := &IntHeap{}
	*h = append(*h, arr[:chunk]...)
	heap.Init(h)
	for i := chunk; i < len(arr); i++ {
		val := heap.Pop(h).(int)
		result = append(result, val)
		heap.Push(h, arr[i])
	}
	for h.Len() > 0 {
		val := heap.Pop(h).(int)
		result = append(result, val)
	}
	return result
}
