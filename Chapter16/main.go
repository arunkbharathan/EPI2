package main

import "fmt"

func main() {
	// towerOfHanoi()
	// nQueens()
	printPermutations()
}

func printPermutations() {
	arr := []int{2, 3, 5, 7}
	result := generatePermutations(arr)
	for _, val := range result {
		fmt.Println(val)
	}
}

// func nQueens() {
// 	size := 2.0
// 	solveNQueens(int(math.Exp2(size)))
// }

// func towerOfHanoi() {
// 	disks := 3
// 	solveTowerOfHanoi(disks)
// }
