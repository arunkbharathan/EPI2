package main

import (
	"container/list"
	"strings"
)

func findSmallestSubarray2(doc string, keyword []string) *[2]int {
	dll := list.New()
	stringListMap := map[string]*list.Element{}
	keywordsLookUp := map[string]bool{}
	var start, end int
	nearestDistance := len(doc)
	for _, val := range keyword {
		keywordsLookUp[val] = true
	}
	for i, word := range strings.Split(strings.ToLower(doc), " ") {
		if _, ok := keywordsLookUp[word]; ok {
			if elem, ok := stringListMap[word]; ok {
				tmp := dll.Remove(elem)
				_ = tmp
				tmp = dll.PushBack(i)
				stringListMap[word] = tmp.(*list.Element)
			} else {
				ele := dll.PushBack(i)
				stringListMap[word] = ele
			}
			if dll.Len() == len(keywordsLookUp) {
				a := dll.Front().Value.(int)
				b := dll.Back().Value.(int)
				if b-a < nearestDistance {
					nearestDistance = b - a
					start = a
					end = b
				}

			}
		}

	}
	return &[2]int{start, end}
}

func findSmallestSubarray1(doc string, keyword []string) *[2]int {
	strArr := strings.Split(strings.ToLower(doc), " ")
	start := len(doc)
	end := len(doc)
	minDistance := len(doc)
	keywordIndex := map[string]int{}
	keywordCount := map[string]int{}
	for _, val := range keyword {
		keywordCount[val] = 0
	}
	var i, j int
	for i < len(strArr)-1 {

		if keywordsPresent(keywordCount) == true {
			i++
			a := minInMap(keywordIndex)
			b := maxInMap(keywordIndex)
			if (b - a) < minDistance {
				minDistance = b - a
				start = a
				end = b
			}
			if _, ok := keywordCount[strArr[i]]; ok {
				keywordCount[strArr[i]]--
			}
		} else {
			if j < len(strArr)-1 {
				j++
				if _, ok := keywordCount[strArr[j]]; ok {
					keywordCount[strArr[j]]++
					keywordIndex[strArr[j]] = j
				}
			} else {
				i++
				a := minInMap(keywordIndex)
				b := maxInMap(keywordIndex)
				if (b - a) < minDistance {
					minDistance = b - a
					start = a
					end = b
				}
				if _, ok := keywordCount[strArr[i]]; ok {
					keywordCount[strArr[i]]--
				}
			}
		}
	}
	return &[2]int{start, end}
}

func keywordsPresent(keywordPresent map[string]int) bool {
	for _, val := range keywordPresent {
		if val == 0 {
			return false
		}
	}
	return true
}

func minInMap(keywordCount map[string]int) int {
	min := 98358959859359
	for _, val := range keywordCount {
		if val < min {
			min = val
		}
	}
	return min
}

func maxInMap(keywordCount map[string]int) int {
	max := -1
	for _, val := range keywordCount {
		if val > max {
			max = val
		}
	}
	return max
}
