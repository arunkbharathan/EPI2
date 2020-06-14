package main

func squareRootOf(num, percision float64) (float64, int) {
	var L, U, M float64
	L = 1.0
	U = num
	if num < 1.0 {
		L = num
		U = 1.0
	}
	numIterations := 0
	for L <= U {
		numIterations++
		M = L + (U-L)*0.5
		squared := M * M
		if absDiff(squared, num) < percision {
			return M, numIterations
		}
		if squared-num > 0 {
			U = M
		} else {
			L = M
		}
	}
	return L, numIterations
}

func absDiff(a, b float64) float64 {
	if a < b {
		return b - a
	}
	return a - b
}
