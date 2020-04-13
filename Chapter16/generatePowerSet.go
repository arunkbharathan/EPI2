package main

import "fmt"

// https://www.geeksforgeeks.org/recursive-program-to-generate-power-set/
func powerSet(arr []int) {
	currSet := []int{}
	generatePowerSet(arr, 0, currSet)
}

func generatePowerSet(arr []int, ind int, currSet []int) {
	if ind == len(arr) {
		fmt.Println(currSet)
		return
	}
	currSet2 := make([]int, 0, len(arr))
	currSet2 = append(currSet2, currSet...)
	currSet2 = append(currSet2, arr[ind])
	generatePowerSet(arr, ind+1, currSet2)
	generatePowerSet(arr, ind+1, currSet)

}
