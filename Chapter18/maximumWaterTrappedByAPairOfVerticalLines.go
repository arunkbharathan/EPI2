package main

func maxWaterTappedBwVerticalLines(wall []int) (int, int, int) {
	i := 0
	j := len(wall) - 1
	lIndex, rIndex, maxArea := -1, len(wall), 0
	for i < j {
		area := 0
		if wall[i] < wall[j] {
			area = (j - i) * wall[i]
		} else {
			area = (j - i) * wall[j]
		}
		if area > maxArea {
			maxArea = area
			rIndex = j
			lIndex = i
		}

		if wall[i] == wall[j] {
			i++
			j--
		} else if wall[i] < wall[j] {
			i++
		} else {
			j--
		}

	}
	return maxArea, lIndex, rIndex
}

func min(a, c int) int {
	if a < c {
		return a
	}
	return c
}
