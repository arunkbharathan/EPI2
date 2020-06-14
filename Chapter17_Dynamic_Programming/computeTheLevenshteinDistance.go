package main

// https://www.geeksforgeeks.org/edit-distance-dp-5/
func computeTheLevenshteinDistance(str1, str2 string) int {
	result := [][]int{}
	distance := -1

	computeEditDistance(str1, str2, &result)
	distance = result[len(result)-1][len(result[len(result)-1])-1]

	return distance
}

func computeEditDistance(str1, str2 string, result *[][]int) {
	for i := 0; i <= len(str1); i++ {
		(*result) = append(*result, make([]int, 0, len(str1)+1))
		for j := 0; j <= len(str2); j++ {
			if i == 0 {
				// (*result)[i][j] = j
				(*result)[i] = append((*result)[i], j)
			} else if j == 0 {
				// (*result)[i][j] = i
				(*result)[i] = append((*result)[i], i)
			} else if str1[i-1] == str2[j-1] {
				// (*result)[i][j] = (*result)[i-1][j-1]
				(*result)[i] = append((*result)[i], (*result)[i-1][j-1])
			} else {
				insert := (*result)[i][j-1]
				delete := (*result)[i-1][j]
				substitue := (*result)[i-1][j-1]
				// (*result)[i][j] = min(insert, delete, substitue)
				(*result)[i] = append((*result)[i], 1+min(insert, delete, substitue))
			}
		}
	}
}

func min(ints ...int) int {
	minVal := 1000000
	for _, v := range ints {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
