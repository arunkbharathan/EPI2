package main

func main() {
	// towerOfHanoi()
	// nQueens()
	// printPermutations()
	printPowerset()
}

func printPowerset() {
	arr := []int{0, 1, 2}
	powerSet(arr)
}

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
