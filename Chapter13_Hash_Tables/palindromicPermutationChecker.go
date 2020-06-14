package main

func testForPalindromicPermutation(arr []string) []bool {

	result := []bool{}
	for _, str := range arr {
		counter := map[rune]int{}
		for _, item := range str {

			if _, ok := counter[item]; ok {
				delete(counter, item)
			} else {
				counter[item]++
			}
		}

		if len(counter) == 1 || len(counter) == 0 {
			result = append(result, true)
		} else {
			result = append(result, false)
		}

	}
	return result
}
