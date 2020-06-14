package main

func searchForFirstK(arr []int, key int) int {
	L := 0
	U := len(arr) - 1
	M := -1
	minM := len(arr)
	for L <= U {
		M = L + (U-L)/2
		if key == arr[M] {
			minM = M
			L = 0
			U = M - 1
		}
		if key > arr[M] {
			L = M + 1
		} else {
			U = M - 1
		}
	}
	return minM
}
