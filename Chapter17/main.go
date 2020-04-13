package main

import "fmt"

func main() {
	countNumberOfScoreCombinations()
}

func countNumberOfScoreCombinations() {
	W := []int{1, 2, 3, 4, 6}
	score := 24
	result := scoreCombinationsFor(W, score)
	fmt.Printf("Number of ways %d can be achieved with scores %v is %d\n", score, W, result)
}
