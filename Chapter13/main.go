package main

import "fmt"

func main() {
	// testPalindromicPermutation()
	// isAnonymousLetterConstrutible()
	// isbnCacheImplementation()
	// lcaOptimized()
	// findKMostFrequentQueries()
	// findNearestRepeatedEntriesInArray()
	findSmallestSubarrayCoveringAllValues()
}

func findSmallestSubarrayCoveringAllValues() {
	excerpt1 := "My paramount object in this struggle is to save the Union , and is not either to save or to destroy slavery. If I could save the Union without freeing any slave, I would do it; and if I could save it by freeing some and leaving others alone, I would also do that."
	keywords1 := []string{"save", "union"}
	excerpt2 := "apple banana banana apple dog cat apple dog apple dog banana cat"
	keywords2 := []string{"banana", "cat"}
	fmt.Println("Method-1")
	arr := findSmallestSubarray1(excerpt1, keywords1)
	fmt.Printf("start: %d end:%d\n", arr[0], arr[1])
	arr = findSmallestSubarray1(excerpt2, keywords2)
	fmt.Printf("start: %d end:%d\n", arr[0], arr[1])
	fmt.Println("Method-2")
	arr = findSmallestSubarray2(excerpt1, keywords1)
	fmt.Printf("start: %d end:%d\n", arr[0], arr[1])
	arr = findSmallestSubarray2(excerpt2, keywords2)
	fmt.Printf("start: %d end:%d\n", arr[0], arr[1])
}

// func findNearestRepeatedEntriesInArray() {
// 	arr := []string{"All", "work", "and", "no", "play", "makes", "for",
// 		"no", "work", "no", "fun", "and", "no", "results"}
// 	nearestRepeatedDistance := nearestRepeatedEntriesIn(arr)
// 	fmt.Println(nearestRepeatedDistance)
// }

// func findKMostFrequentQueries() {
// 	for _, val := range findFrequentQueries(10) {
// 		fmt.Println(*val)
// 	}
// }

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
