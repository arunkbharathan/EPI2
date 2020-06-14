package main

func scoreCombinationsFor(W []int, score int) int {
	resultSlice := [][]int{}
	for i, point := range W {
		row := make([]int, 1, score+1)
		// A[0][0] = 1
		row[0] = 1
		resultSlice = append(resultSlice, row)
		for j := 1; j <= score; j++ {
			var withOutPlay, withPlay int
			if i >= 1 {
				withOutPlay = resultSlice[i-1][j]
			} else {
				withOutPlay = 0
			}
			if j < point {
				withPlay = 0
			} else {
				withPlay = resultSlice[i][j-point]
			}
			resultSlice[i] = append(resultSlice[i], withOutPlay+withPlay)
		}
	}
	return resultSlice[len(resultSlice)-1][len(resultSlice[len(resultSlice)-1])-1]
}
