package main

import (
	"sort"
)

func threeSum(arr []int, num int) []int {
	sort.Sort(sort.IntSlice(arr))
	result := []int{}

	for _, k := range arr {

		balanceSum := num - k
		result = twoSumProblem(arr, balanceSum)
		if len(result) > 0 {
			result = append([]int{k}, result...)
			break
		}
	}

	return result
}

func twoSumProblem(arr []int, num int) []int {
	result := []int{}
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]+arr[j] > num {
			j--
		} else if arr[i]+arr[j] < num {
			i++
		} else {
			result = append(result, arr[i], arr[j])
			break
		}
	}
	return result
}
