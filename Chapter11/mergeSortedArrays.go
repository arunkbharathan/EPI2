package main

import "container/heap"

func mergeSortedArrays(arrays [][]int) []int {
	arrMap := map[int]int{}
	result := []int{}
	h := &IntHeap{}
	heap.Init(h)
	for ind, array := range arrays {
		if len(array) > 0 {
			arrMap[array[0]] = ind
			heap.Push(h, array[0])
		}
	}

	for h.Len() > 0 {

		minVal := heap.Pop(h).(int)
		result = append(result, minVal)
		minValIndex := arrMap[minVal]
		arrays[minValIndex] = arrays[minValIndex][1:]
		if len(arrays[minValIndex]) > 0 {
			heap.Push(h, arrays[minValIndex][0])
		}
		arrMap = map[int]int{}
		for ind, array := range arrays {
			if len(array) > 0 {
				arrMap[array[0]] = ind
			}
		}

	}

	return result
}
