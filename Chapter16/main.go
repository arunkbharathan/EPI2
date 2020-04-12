package main

import "math"

func main() {
	// towerOfHanoi()
	nQueens()
}
func nQueens() {
	size := 2.0
	solveNQueens(int(math.Exp2(size)))
}

// func towerOfHanoi() {
// 	disks := 3
// 	solveTowerOfHanoi(disks)
// }
