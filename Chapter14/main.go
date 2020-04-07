package main

import "fmt"

func main() {
	// findIntersectionOf2SortedArrays()
	mergeTwoSortedArrays()
}

func mergeTwoSortedArrays() {
	arr1 := []int{5, 13, 17}
	arr2 := []int{3, 7, 11, 19}
	arr := merge2SortedArrays(arr1, arr2)
	fmt.Printf("Array1:%v\nArray2:%v\nMerged:%v\n", arr1, arr2, arr)
}

// func findIntersectionOf2SortedArrays() {
// 	arr1 := []int{2, 3, 3, 5, 5, 6, 7, 7, 8, 12}
// 	arr2 := []int{5, 5, 6, 8, 8, 9, 10, 10}
// 	arr := findIntersection(arr1, arr2)
// 	fmt.Printf("Array1:%v\nArray2:%v\nIntersection:%v\n", arr1, arr2, arr)
// 	fmt.Println()
// 	arr1 = []int{2, 3, 3, 5, 7, 11}
// 	arr2 = []int{3, 3, 7, 15, 31}
// 	arr = findIntersection(arr1, arr2)
// 	fmt.Printf("Array1:%v\nArray2:%v\nResult:%v\n", arr1, arr2, arr)
// }
