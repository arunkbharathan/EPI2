package main

func merge2SortedArrays(arr1, arr2 []int) []int {
	len1 := len(arr1)
	len2 := len(arr2)
	for range arr2 {
		arr1 = append(arr1, 0)
	}
	m := len1 - 1
	n := len2 - 1
	writeIndex := len1 + len2 - 1
	for m >= 0 && n >= 0 {
		if arr1[m] < arr2[n] {
			arr1[writeIndex] = arr2[n]
			n--
			writeIndex--

		} else if arr1[m] >= arr2[n] {
			arr1[writeIndex] = arr1[m]
			m--
			writeIndex--
		}

	}

	for ; m >= 0; m-- {
		arr1[writeIndex] = arr1[m]
		writeIndex--
	}
	for ; n >= 0; n-- {
		arr1[writeIndex] = arr2[n]
		writeIndex--
	}

	return arr1
}
