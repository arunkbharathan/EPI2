package main

func computeEnclosedRegions(region [][]int, num int) [][2]int {
	visited := make([][]bool, len(region))
	result := [][2]int{}
	for i := range visited {
		visited[i] = make([]bool, len(region[0]))
	}

	queue := make([][2]int, 0, len(region))

	for i := 0; i < len(region); i++ {
		for j := 0; j < len(region[0]); j++ {
			if i == 0 || j == 0 || i == len(region)-1 || j == len(region[0])-1 {
				if region[i][j] == num {
					queue = append(queue, [2]int{i, j})
				}
			}
			if region[i][j] == 1-num {
				visited[i][j] = true
			}
		}
	}
	directions := [][2]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node[0]][node[1]] = true
		for _, dirVector := range directions {
			next := [2]int{node[0] + dirVector[0], node[1] + dirVector[1]}
			if next[0] < 0 || next[0] >= len(region) || next[1] < 0 || next[1] >= len(region[0]) || visited[next[0]][next[1]] == true {
				continue
			}
			queue = append(queue, next)
		}
	}

	for i := range visited {
		for j := range visited[i] {
			if !visited[i][j] {
				result = append(result, [2]int{i, j})

			}
		}
	}

	return result
}
