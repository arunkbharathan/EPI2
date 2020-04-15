package main

import "fmt"

func main() {
	// countNumberOfScoreCombinations()
	// editDistance()
	// traverse2DArray()
	// searchForASequenceIn2DArray()
	solveKnapSackProblem()
}

func solveKnapSackProblem() {
	W := []int{5, 3, 2, 4}
	P := []int{60, 50, 30, 70}
	weightConstraint := 5
	maxVal, items := knapSackMaximumValue(W, P, weightConstraint)
	fmt.Printf("Maximum profit for the given knapsack weight is %d\nFor profits: %v and weights: %v and items are %v + 1\n", maxVal, P, W, items)
}

// func searchForASequenceIn2DArray() {
// 	seq := []int{1, 3, 4, 6}
// 	grid := [][]int{
// 		{1, 2, 3},
// 		{3, 4, 5},
// 		{5, 6, 7},
// 	}

// 	fmt.Printf("%v occurs in %v :%t\n", seq, grid, search2DArray(seq, grid))

// }

// func traverse2DArray() {
// 	arraySize := 5
// 	ways := computeWaysToTraverseA2DArray(arraySize)
// 	fmt.Println(ways, "ways are there.")
// }

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
