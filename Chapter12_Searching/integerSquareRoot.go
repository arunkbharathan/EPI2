package main

func integerSquareRootOf(num int) int {
	L := 0
	U := num
	for L <= U {
		M := L + (U-L)/2
		square := M * M
		if square == num {
			return M
		}
		if square > num {
			U = M - 1
		}
		if square < num {
			L = M + 1
		}
	}
	return L - 1
}
