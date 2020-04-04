package main

import "fmt"

func main() {
	// testPalindromicPermutation()
	isAnonymousLetterConstrutible()
}

func isAnonymousLetterConstrutible() {
	fmt.Println(testForAnonymousLetterConstrutible("./magazine", "./letter"))
}

// func testPalindromicPermutation() {
// 	arr := []string{"level", "rotator", "foobaraboof", "edified", "corona", "virus", "ipopoi"}
// 	result := testForPalindromicPermutation(arr)
// 	fmt.Println(arr)
// 	fmt.Println(result)
// }
