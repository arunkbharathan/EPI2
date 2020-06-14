package main

import "fmt"

func spiralize(arr [][]int, size int, start int) {

	if size == 1 {
		fmt.Println(arr[start+size-1][start+size-1])
		return
	}
	if size <= 0 {
		return
	}

	for i := start; i < start+size; i++ {
		fmt.Print(arr[start][i], " ")

	}
	for i := start + 1; i < start+size-1; i++ {
		fmt.Printf("\n   %d", arr[i][start+size-1])
	}
	fmt.Println()
	for i := start + size - 1; i >= start; i-- {
		fmt.Print(arr[start+size-1][i], " ")
	}
	fmt.Println()
	for i := start + size - 2; i > start; i-- {
		fmt.Printf("%d   \n", arr[i][start])
	}
	start++
	spiralize(arr, size-2, start)
}
