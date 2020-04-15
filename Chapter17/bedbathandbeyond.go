package main

func isWordInDictionary(word string, dictionery map[string]int) bool {
	result := make([][]int, len(word))
	// splitPoint := make([][]int, len(word)*len(word))

	for i := range result {
		result[i] = make([]int, len(word))
		for j := range result[i] {
			if i == j {
				letter := string(word[i])
				if _, ok := dictionery[letter]; ok {
					result[i][j] = 1
				} else {
					result[i][j] = 0
				}
			} else {
				result[i][j] = -1
			}
			result[i][j] = -1
		}
	}

	checkForWordInDictionary(&result, word, 0, len(word), &dictionery)
	return false
}

func checkForWordInDictionary(result *[][]int, word string, n, k int, dictionery *map[string]int) int {
	array := *result
	dict := *dictionery
	// if n > k {
	// 	return 0
	// }
	if len(word) == 0 {
		return 1
	}
	if _, ok := dict[word]; ok {
		array[n][k] = 1
		return 1
	}

	if array[n][k] == -1 {

		for i := 0; i < k; i++ {
			for j := 0; j < i; j++ {
				word1, word2 := word[0:i], word[i:k+1]
				_, _ = word1, word2
				val1 := checkForWordInDictionary(result, word, 0, i, dictionery)
				val2 := checkForWordInDictionary(result, word, i, k+1, dictionery)
				if val1 == 1 && val2 == 1 {
					array[n][k] = 1
				} else {
					array[n][k] = 0
				}
			}
		}

	}
	return array[n][k]

}
