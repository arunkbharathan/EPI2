package main

func main() {
	// towerOfHanoi()
	// nQueens()
	// printPermutations()
	// printPowerset()
	solveSudoku()
}

func solveSudoku() {
	arr := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	sudokuSolver(arr)
}

// func printPowerset() {
// 	arr := []int{0, 1, 2}
// 	powerSet(arr)
// }

// func printPermutations() {
// 	arr := []int{2, 3, 5, 7}
// 	result := generatePermutations(arr)
// 	for _, val := range result {
// 		fmt.Println(val)
// 	}
// }

// func nQueens() {
// 	size := 2.0
// 	solveNQueens(int(math.Exp2(size)))
// }

// func towerOfHanoi() {
// 	disks := 3
// 	solveTowerOfHanoi(disks)
// }
