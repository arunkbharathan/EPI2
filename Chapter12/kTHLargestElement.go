package main

import (
	"math/rand"
	"time"
)

func kTHLargestElement(arr []int, k int) int {
	L := 0
	R := len(arr) - 1
	kthLargestProperIndex := len(arr) - k
	for {
		pivotIndex := randIntBetweenAndIncluding(L, R)
		pivotNewIndex := partitionArray(arr, pivotIndex)
		if kthLargestProperIndex == pivotNewIndex {
			return arr[pivotNewIndex]
		} else if pivotNewIndex < kthLargestProperIndex {
			L = pivotNewIndex + 1
		} else {
			R = pivotNewIndex - 1
		}

	}
	return 0
}

func partitionArray(arr []int, index int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[index] && i < index {
			arr[i], arr[index] = arr[index], arr[i]
			index = i
			i = min(index, i)
		}
		if arr[i] < arr[index] && i > index {
			arr[i], arr[index] = arr[index], arr[i]
			index = i
			i = min(index, i)
		}
	}
	return index
}

func randIntBetweenAndIncluding(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	min := a
	max := b
	return (rand.Intn(max-min+1) + min)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
