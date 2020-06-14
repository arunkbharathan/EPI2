package main

func findIntersection(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0
	ffa, ffb := false, false
	for i < len(arr1) && j < len(arr2) {

		if arr1[i] == arr2[j] {
			if len(result) > 0 {
				if result[len(result)-1] != arr1[i] {
					result = append(result, arr1[i])
				}
			} else {
				result = append(result, arr1[i])
			}
			ffa, ffb = true, true
		}
		if arr1[i] < arr2[j] {
			ffa = true
			ffb = false
		}
		if arr1[i] > arr2[j] {
			ffa = false
			ffb = true
		}
		if ffa {
			i++
		}
		if ffb {
			j++
		}
	}
	return result
}
