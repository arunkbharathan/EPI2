package main

func smallestInCyclicallySorted(arr []int) int {

	result := divideNConquer(arr, 0, len(arr)-1)
	return result
}

func divideNConquer(arr []int, L, U int) int {

	if (U - L) == 1 {
		return min(arr[L], arr[U], L, U)
	}
	if (U - L) == 0 {
		return L
	}
	M := L + (U-L)/2
	indA := divideNConquer(arr, L, M)
	indB := divideNConquer(arr, M+1, U)
	minIndex := min(arr[indA], arr[indB], indA, indB)

	return minIndex
}

func min(a, b, l, m int) int {
	if a < b {
		return l
	}
	return m
}
