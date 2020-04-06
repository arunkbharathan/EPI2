package main

func nearestRepeatedEntriesIn(arr []string) int {
	repeationDistance := len(arr)
	repeatWordHolder := map[string]int{}
	for i, val := range arr {
		if ind, ok := repeatWordHolder[val]; ok {
			repeationDistance = min(i-ind, repeationDistance)
			repeatWordHolder[val] = i
		} else {
			repeatWordHolder[val] = i
		}
	}
	return repeationDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
