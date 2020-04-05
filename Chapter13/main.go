package main

import "fmt"

func main() {
	// testPalindromicPermutation()
	// isAnonymousLetterConstrutible()
	// isbnCacheImplementation()
	// lcaOptimized()
	findKMostFrequentQueries()
}

func findKMostFrequentQueries() {
	for _, val := range findFrequentQueries(10) {
		fmt.Println(*val)
	}
}

// func lcaOptimized() {
// 	bt1, bt2 := btGetTwoNodes()
// 	fmt.Println(findLCAOptimized(bt1, bt2).data)
// }

// func isbnCacheImplementation() {
// 	isbnBookCache()
// }

// func isAnonymousLetterConstrutible() {
// 	fmt.Println(testForAnonymousLetterConstrutible("./magazine", "./letter"))
// }

// func testPalindromicPermutation() {
// 	arr := []string{"level", "rotator", "foobaraboof", "edified", "corona", "virus", "ipopoi"}
// 	result := testForPalindromicPermutation(arr)
// 	fmt.Println(arr)
// 	fmt.Println(result)
// }
