package main

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
)

func findNclosestStars(n int) []int {
	h := &IntMaxHeap{}
	result := []int{}
	heap.Init(h)

	fs, err := os.Open("./stardataset")
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()
	fileScanner := bufio.NewScanner(fs)
	for fileScanner.Scan() {
		intVal, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Println(err)
			continue
		}
		heap.Push(h, intVal)
		if h.Len() > n {
			heap.Pop(h)
		}
		// fmt.Println(intVal)
	}
	for h.Len() > 0 {
		result = append(result, h.Pop().(int))
	}
	return result
}
