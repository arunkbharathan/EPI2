package main

import (
	"math/rand"
	"time"
)

func randomizeArray(arr [10]int, size int) [10]int {
	count := 0
	for i := size; i > 0; i-- {
		randIndex := randInt(0, i)

		arr[count], arr[randIndex] = arr[randIndex], arr[count]
		count++
	}
	return arr
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
