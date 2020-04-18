package main

func paintBooleanMatrix(boolMatrix [][]uint, paint [2]int) [][2]int {
	colorToPaint := 1 - boolMatrix[paint[0]][paint[1]]
	news := [][2]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	quu := [][2]int{}
	result := [][2]int{}

	quu = append(quu, paint)

	for len(quu) > 0 {
		nodePoint := quu[0]
		quu = quu[1:]
		result = append(result, nodePoint)

		boolMatrix[nodePoint[0]][nodePoint[1]] = colorToPaint

		for _, dirVector := range news {
			next := [2]int{nodePoint[0] + dirVector[0], nodePoint[1] + dirVector[1]}
			if next[0] < 0 || next[0] >= len(boolMatrix) || next[1] < 0 || next[0] >= len(boolMatrix[0]) || boolMatrix[next[0]][next[1]] == colorToPaint {
				continue
			}
			quu = append(quu, next)
		}

	}
	return result
}
