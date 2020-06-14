package main

func searchMazeForExit(maze [][]int, start, end [2]int) [][2]int {
	path := make([][2]int, 0, len(maze))
	maze[start[0]][start[1]] = 1
	path = append(path, start)
	if !searchMazeHelper(&maze, &path, start, end) {
		path = path[1:]
	}
	return (path)
}

func searchMazeHelper(forest *[][]int, way *[][2]int, start, end [2]int) bool {
	maze := *forest
	path := *way
	kMove := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	if start == end {
		return true
	}
	// curr := start
	for _, newMove := range kMove {
		curr := path[len(path)-1]
		next := [2]int{curr[0] + newMove[0], curr[1] + newMove[1]}

		if isFeaseble(forest, next) {
			maze[next[0]][next[1]] = 1
			*way = append(path, next)
			if searchMazeHelper(forest, way, next, end) {
				return true
			}
		}
		if len(*way) != 1 {
			*way = (*way)[1:]
		}
	}
	return false
}

// 0 => no obstacle
// 1 => obstacle
func isFeaseble(forest *[][]int, locs [2]int) bool {
	x := locs[0]
	y := locs[1]
	maze := *forest
	if x < 0 || y < 0 || y >= len(maze) || x >= len(maze[0]) || maze[x][y] == 1 {
		return false
	}
	return true
}
