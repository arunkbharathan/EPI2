package main

func isValidSudoku(arr [9][9]int) bool {

	for i := 0; i < len(arr); i++ {
		existsArray := make([]bool, 10)
		for j := 0; j < len(arr); j++ {
			if arr[i][j] != 0 {
				if existsArray[arr[i][j]] {
					return false
				}
				existsArray[arr[i][j]] = true

			}
		}
	}

	for j := 0; j < len(arr); j++ {
		existsArray := make([]bool, 10)
		for i := 0; i < len(arr); i++ {
			if arr[i][j] != 0 {
				if existsArray[arr[i][j]] {
					return false
				}
				existsArray[arr[i][j]] = true

			}
		}
	}

	for j := 0; j < len(arr); j++ {
		existsArray := make([]bool, 10)
		for i := 0; i < len(arr); i++ {
			if arr[i][j] != 0 {
				if existsArray[arr[i][j]] {
					return false
				}
				existsArray[arr[i][j]] = true

			}
		}
	}

	for l := 0; l < len(arr); l += 3 {
		for m := 0; m < len(arr); m += 3 {
			existsArray := make([]bool, 10)
			for i := l; abs(i-l) < 3; i++ {
				for j := m; abs(j-m) < 3; j++ {
					if arr[i][j] != 0 {
						if existsArray[arr[i][j]] {
							return false
						}
						existsArray[arr[i][j]] = true

					}
				}
			}
		}
	}

	return true
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
