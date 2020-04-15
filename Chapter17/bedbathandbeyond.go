package main

func isWordInDictionary(word string, dictionery map[string]int) []string {
	lastLength := make([]int, len(word))

	for i := range lastLength {
		lastLength[i] = -1
	}
	for i := 0; i < len(word); i++ {
		str := word[0 : i+1]
		if _, ok := dictionery[str]; ok {
			lastLength[i] = i + 1
		}

		if lastLength[i] == -1 {
			for j := 0; j < i; j++ {
				if lastLength[j] != -1 {
					str := word[j+1 : i+1]
					if _, ok := dictionery[str]; ok {
						lastLength[i] = i - j
						break
					}
				}
			}
		}
	}

	decompositions := []string{}

	if lastLength[len(lastLength)-1] != -1 {
		idx := len(word) - 1

		for idx >= 0 {
			str := word[idx+1-lastLength[idx] : idx+1]
			decompositions = append(decompositions, str)
			idx -= lastLength[idx]
		}
		for i := 0; i < len(decompositions)/2; i++ {
			j := len(decompositions) - 1 - i
			decompositions[i], decompositions[j] = decompositions[j], decompositions[i]
		}
	}

	return decompositions
}
