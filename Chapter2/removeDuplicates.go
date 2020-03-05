package main

func removeDuplicates(arr [10]int) [10]int {
	newi := 0
	dupdetect := false
	for i := 1; i < len(arr); i++ {
		if dupdetect {
			if arr[i] != arr[newi] {
				newi++
				arr[newi] = arr[i]
			}
			continue
		}

		if arr[i] != arr[i-1] {
			newi++
		} else {
			dupdetect = true
		}
	}
	if dupdetect {

		for i := newi + 1; i < len(arr); i++ {
			arr[i] = 0
		}
	}
	return arr
}
