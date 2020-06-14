package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

func streamingMedian() {
	maxH := &FloatMaxHeap{}
	minH := &FloatHeap{}
	heap.Init(maxH)
	heap.Init(minH)

	// fs, err := os.Open("./numbers")
	// fs, err := os.Open("./sortedNumbers")
	fs, err := os.Open(os.Stdin.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()
	fileScanner := bufio.NewScanner(fs)
	median := 0.0
	for fileScanner.Scan() {
		intVal, err := strconv.Atoi(fileScanner.Text())
		floatVal := float64(intVal)
		if err != nil {
			log.Println(err)
			continue
		}
		// Push first to minheap
		if minH.Len() == 0 {
			heap.Push(minH, floatVal)
		} else {
			// Push based on top of minHeap(maxHeap is also fine)
			minHVal := heap.Pop(minH).(float64)
			if floatVal >= minHVal {
				heap.Push(minH, floatVal)
			} else {
				heap.Push(maxH, floatVal)
			}
			heap.Push(minH, minHVal)
		}
		// Ensure min_heap and max_heap have equal number of elements if
		// an even number of elements is read; otherwise, min_heap must have
		// one more element than max_heap.

		if minH.Len() > maxH.Len()+1 {
			heap.Push(maxH, (heap.Pop(minH)))
		} else if maxH.Len() > minH.Len() {
			heap.Push(minH, heap.Pop(maxH))
		}

		//Calculate median
		if maxH.Len() == minH.Len() {
			rVal := heap.Pop(minH).(float64)
			lVal := heap.Pop(maxH).(float64)
			fmt.Println((rVal + lVal) / 2)
			heap.Push(minH, rVal)
			heap.Push(maxH, lVal)
		} else {
			median = heap.Pop(minH).(float64)
			fmt.Println(median)
			heap.Push(minH, median)
		}
	}
	// fmt.Println(minH.Len())
	// fmt.Println(maxH.Len())
}
