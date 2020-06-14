package main

func generatePermutations(arr []int) [][]int {
	result := make([][]int, 0, 30)
	directedPermutation(0, arr, &result)
	return result
}

func directedPermutation(i int, arr []int, result *[][]int) {
	if i == len(arr) {
		intSlice := make([]int, 0, 4)
		intSlice = append(intSlice, arr...)
		*result = append(*result, intSlice)
		return
	}

	for j := i; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		directedPermutation(i+1, arr, result)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
