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

	fs, err := os.Open("./numbers")
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()
	fileScanner := bufio.NewScanner(fs)
	cnt := 0
	preMedian := 0.0
	median := 0.0
	for fileScanner.Scan() {
		cnt++
		intVal, err := strconv.Atoi(fileScanner.Text())
		floatVal := float64(intVal)
		if err != nil {
			log.Println(err)
			continue
		}

		if cnt == 1 {
			median = floatVal
			fmt.Println(median)
			preMedian = median
		} else if cnt == 2 {
			median = (preMedian + floatVal) / 2
			fmt.Println(median)
			if floatVal < median {
				heap.Push(maxH, floatVal)
				heap.Push(minH, preMedian)
			} else {
				heap.Push(minH, floatVal)
				heap.Push(maxH, preMedian)
			}
			preMedian = median
		} else {
			//odd cnt means a number in array is middle element
			// else the avg of mid 2 elements split sorted array to 2 equal half
			if cnt%2 == 0 {
				if floatVal > preMedian {

					heap.Push(minH, floatVal)

					rVal := heap.Pop(minH).(float64)
					median = (rVal + preMedian) / 2
					if rVal < median {
						heap.Push(maxH, rVal)
					} else {
						heap.Push(minH, rVal)
					}
					if preMedian < median {
						heap.Push(maxH, preMedian)
					} else {
						heap.Push(minH, preMedian)
					}

				}
				if floatVal < preMedian {
					heap.Push(maxH, floatVal)
					lVal := heap.Pop(maxH).(float64)
					median = (lVal + preMedian) / 2
					if lVal < median {
						heap.Push(maxH, lVal)
					} else {
						heap.Push(minH, lVal)
					}

				}
				if floatVal == median {
					heap.Push(minH, floatVal)
				}
				fmt.Println(median)
				preMedian = median
			} else {
				if floatVal > preMedian {

					heap.Push(minH, floatVal)

					rVal := heap.Pop(minH).(float64)
					median = rVal

				}
				if floatVal < median {
					heap.Push(maxH, floatVal)

					lVal := heap.Pop(maxH).(float64)
					median = lVal

				}
				if floatVal == median {
					heap.Push(minH, floatVal)
				}
				fmt.Println(median)
				preMedian = median
			}

		}
	}
	fmt.Println(maxH.Len())
	fmt.Println(maxH.Len())
}
