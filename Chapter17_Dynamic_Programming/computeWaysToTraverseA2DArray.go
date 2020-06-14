package main

func computeWaysToTraverseA2DArray(arrSize int) int {
	resultSlice := [][]int{}
	traverseA2DArray(&resultSlice, arrSize)
	ways := resultSlice[len(resultSlice)-1][len(resultSlice[len(resultSlice)-1])-1]
	return ways
}

func traverseA2DArray(result *[][]int, size int) {
	for i := 0; i < size; i++ {
		*result = append(*result, make([]int, 0, size))
		for j := 0; j < size; j++ {
			if i == 0 && j == 0 {
				(*result)[i] = append((*result)[i], 0)
			} else if i == 0 {
				(*result)[i] = append((*result)[i], 1)
			} else if j == 0 {
				// (*result)[i][j] = 1
				(*result)[i] = append((*result)[i], 1)
			} else {
				// (*result)[i][j] = (*result)[i-1][j] + (*result)[i][j-1]
				(*result)[i] = append((*result)[i], (*result)[i-1][j]+(*result)[i][j-1])
			}
		}
	}
}
