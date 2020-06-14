package main

func smallestInCyclicallySorted(arr []int) int {
	L := 0
	R := len(arr) - 1
	for L < R {
		M := L + (R-L)/2
		if arr[M] < arr[R] {
			R = M
		} else {
			L = M + 1
		}
	}
	return L
}
