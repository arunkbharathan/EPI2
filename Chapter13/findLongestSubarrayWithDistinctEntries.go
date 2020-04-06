package main

import "strings"

func longestSubarrayWithDistinctEntries(arr string) int {
	stringLookUp := map[string]int{}
	longestSubsequence := -1
	startIndex := 0
	strArray := strings.Split(strings.ToLower(arr), " ")
	for i, str := range strArray {
		ind, ok := stringLookUp[str]
		if ind >= startIndex && ok {
			longestSubsequence = max(i-startIndex, longestSubsequence)
			startIndex = ind + 1
			stringLookUp[str] = i
		} else {
			stringLookUp[str] = i
		}
	}
	longestSubsequence = max(len(strArray)-startIndex, longestSubsequence)
	return longestSubsequence
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
