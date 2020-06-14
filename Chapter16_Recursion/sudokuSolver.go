package main

import "fmt"

const emptyEntry = 0

func sudokuSolver(partialGrid [][]int) {
	solvePartialSudoku(0, 0, &partialGrid)
	fmt.Println(partialGrid)
}

func solvePartialSudoku(i, j int, partialGrid *[][]int) bool {
	if i == len(*partialGrid) {
		i = 0
		j++
		if j == len(*partialGrid) {
			return true
		}
	}

	if (*partialGrid)[i][j] != emptyEntry {
		return solvePartialSudoku(i+1, j, partialGrid)
	}

	for val := 1; val <= len(*partialGrid); val++ {
		if isValidToAddDigitInSudoku(*partialGrid, i, j, val) {
			(*partialGrid)[i][j] = val
			if solvePartialSudoku(i+1, j, partialGrid) {
				return true
			}
		}
	}
	(*partialGrid)[i][j] = emptyEntry
	return false
}

func isValidToAddDigitInSudoku(partialGrid [][]int, i, j, val int) bool {

	//Check row
	for k := 0; k < len(partialGrid); k++ {
		if val == partialGrid[k][j] {
			return false
		}
	}

	//Check coloumn
	for k := 0; k < len(partialGrid); k++ {
		if val == partialGrid[i][k] {
			return false
		}
	}

	//Check region constraints
	regionSize := 3
	I := i / regionSize
	J := j / regionSize
	for a := 0; a < regionSize; a++ {
		for b := 0; b < regionSize; b++ {
			if val == partialGrid[regionSize*I+a][regionSize*J+b] {
				return false
			}
		}
	}
	return true
}
