package main

import "fmt"

func main() {
	// countNumberOfScoreCombinations()
	// editDistance()
	traverse2DArray()
}

func traverse2DArray() {
	arraySize := 5
	ways := computeWaysToTraverseA2DArray(arraySize)
	fmt.Println(ways, "ways are there.")
}

// func editDistance() {
// 	str1 := "Carthorse"
// 	str2 := "Orchestra"
// 	dist := computeTheLevenshteinDistance(str1, str2)
// 	fmt.Printf("Levenshtein distance between \"%s\" and \"%s\" is %d\n", str1, str2, dist)
// }

// func countNumberOfScoreCombinations() {
// 	W := []int{1, 2, 3, 4, 6}
// 	score := 24
// 	result := scoreCombinationsFor(W, score)
// 	fmt.Printf("Number of ways %d can be achieved with scores %v is %d\n", score, W, result)
// }
