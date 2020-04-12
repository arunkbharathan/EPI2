package main

import "fmt"

func solveNQueens(size int) [][]int {
	arr := []int{}
	// isQueensPositioningValid(&[]int{1, 3})
	result := nQueensPositioning(arr, size)
	fmt.Println(result)
	return result
}

func nQueensPositioning(arr []int, size int) [][]int {
	result := make([][]int, 0, size)
	if len(arr) == size {
		result = append(result, arr)
		// fmt.Println(result)
	} else {
		for i := 0; i < size; i++ {
			// this solve result getting overwritten on int append in for loop.
			slice := make([]int, 0, size)
			slice = append(slice, arr...)
			slice = append(slice, i)
			if isQueensPositioningValid(slice) {
				result = append(result, nQueensPositioning(slice, size)...)
			}
			arr = slice[:len(slice)-1]

		}
	}

	return result
}

func isQueensPositioningValid(arr []int) bool {
	for i, k := range arr {
		x1, y1 := i, k
		for j := i + 1; j < len(arr); j++ {
			x2, y2 := j, (arr)[j]
			a, b := diffMod(x2, x1), diffMod(y2, y1)
			if (a == 0 && b != 0) || (a != 0 && b == 0) || a == b {
				return false
			}

		}
	}
	return true
}

func diffMod(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
