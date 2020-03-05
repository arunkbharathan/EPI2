package main

import (
	"fmt"
)

func main() {
	// arr := [...]int{0, 2, 0, 1, 2, 1, 2, 1, 0, 0}

	// arr := [...]int{5, 4, 2, 3, 1}
	// pivotVarInd := 3
	// fmt.Println(arr, pivotVarInd)
	// currentPivotInd := pivot1(&arr, pivotVarInd, 0, len(arr)-1)
	// fmt.Println(arr, currentPivotInd)
	// pivotVarInd = currentPivotInd
	// pivot2(&arr, pivotVarInd, currentPivotInd, len(arr)-1)
	// fmt.Println(arr)

	// fmt.Println(maxProfit([...]int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}))

	// arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(randomizeArray(arr, 10))

	// arr := [][]int{
	// 	{1, 2, 3, 4, 5},
	// 	{6, 7, 8, 9, 10},
	// 	{11, 12, 13, 14, 15},
	// 	{16, 17, 18, 19, 20},
	// 	{21, 22, 23, 24, 25},
	// 	// {1, 2, 3, 4},
	// 	// {5, 6, 7, 8},
	// 	// {9, 10, 11, 12},
	// 	// {13, 14, 15, 16},
	// 	// {1, 2, 3},
	// 	// {4, 5, 6},
	// 	// {7, 8, 9},
	// }
	// // fmt.Println(arr)
	// spiralize(arr, len(arr), 0)

	// arr := [4]int{0, 9, 9, 9}
	// fmt.Printf("%v\n%v", arr, arrayIncrement(arr))

	// arr := [...]int{2, 3, 5, 5, 7, 11, 11, 11, 13, 13}
	// fmt.Printf("%v\n%v", arr, removeDuplicates(arr))

	// fmt.Println(primesUpto(1000000))

	arr := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	fmt.Println(isValidSudoku(arr))
}

func pivot1(arr *[10]int, p, start, end int) int {
	for i := end; i >= start; i-- {
		if arr[p] < arr[i] && p > i {
			arr[p], arr[i] = arr[i], arr[p]
			p, i = i, p
		} else if arr[p] > arr[i] && p < i {
			arr[p], arr[i] = arr[i], arr[p]
			p = i
		}

		// fmt.Println(arr)
	}
	return p
}

func pivot2(arr *[10]int, p, start, end int) {
	for i := start; i <= end; i++ {
		if arr[i] != arr[p] {
			arr[i], arr[end] = arr[end], arr[i]
		}
		end--
		fmt.Println(arr)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
