package main

import "fmt"

func main() {
	// searchExample()
	// searchSortedArrayForFirstOccurrencsOfK()
	// smallestInCyclicallySortedArray()
	// findIntegerSquareRoot()
	// findSquareRoot()
	// findKthLargestElementInArray()
	findMissingAndDuplicateInArray()
}

func findMissingAndDuplicateInArray() {
	arr := []int{5, 3, 0, 3, 1, 2}
	m, d := findMissingAndDuplicate(arr)
	fmt.Println(m, d)
}

// func findMissinfIPAddress() {
// 	findMissingIP()
// }

// func findKthLargestElementInArray() {
// 	arr := []int{3, 2, 1, 5, 4, 6, 9, 8, 7, 10, 654, 43, 576, 34, 5757, 534}
// 	k := 10
// 	res := kTHLargestElement(arr, k)
// 	fmt.Println(arr, k)
// 	fmt.Println(res)
// }

// func findSquareRoot() {
// 	var num float64 = 0.089
// 	sr, itr := squareRootOf(num, 0.001)
// 	fmt.Printf("The square root of %.2f is %.2f and its square is %.2f \nand number of iterations it took is %d.\n", num, sr, sr*sr, itr)
// 	num = 924783849738
// 	sr, itr = squareRootOf(num, 0.001)
// 	fmt.Printf("The square root of %.2f is %.2f and its square is %.2f \nand number of iterations it took is %d.\n", num, sr, sr*sr, itr)
// }

// func findIntegerSquareRoot() {
// 	num := 1000
// 	sr := integerSquareRootOf(num)
// 	fmt.Printf("Integer square root of %d is %d and its square is %d.\n", num, sr, sr*sr)
// }

// func smallestInCyclicallySortedArray() {
// 	arr := []int{378, 478, 550, 631, 635, 103, 203, 220, 234, 279}
// 	index := smallestInCyclicallySorted(arr)
// 	fmt.Printf("Smallest element in array %v is at index %d\n", arr, index)
// }

// func searchSortedArrayForFirstOccurrencsOfK() {
// 	arr := []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}
// 	x := 108
// 	index := searchForFirstK(arr, x)
// 	fmt.Printf("First occurrence of %d in %v is at index %d\n", x, arr, index)
// 	arr = []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}
// 	x = 285
// 	index = searchForFirstK(arr, x)
// 	fmt.Printf("First occurrence of %d in %v is at index %d\n", x, arr, index)
// 	arr = []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9, 10, 11, 12}
// 	x = 3
// 	index = searchForFirstK(arr, x)
// 	fmt.Printf("First occurrence of %d in %v is at index %d\n", x, arr, index)
// }

// func searchExample() {
// 	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
// 	x := 6

// 	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
// 	if i < len(a) && a[i] == x {
// 		fmt.Printf("found %d at index %d in %v\n", x, i, a)
// 	} else {
// 		fmt.Printf("%d not found in %v\n", x, a)
// 	}
// }
