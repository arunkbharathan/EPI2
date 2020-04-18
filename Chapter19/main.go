package main

import "fmt"

func main() {
	searchMaze()
}

func searchMaze() {
	// 0 => no obstacle
	// 1 => obstacle
	maze := [][]int{
		{1, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 1, 1, 0, 1, 1},
		{0, 0, 0, 1, 1, 1, 0, 0, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 1, 0, 0, 0},
		{1, 0, 1, 1, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
	}

	start := [2]int{9, 0}
	end := [2]int{0, 9}

	fmt.Printf("Maze\n\n")
	for _, row := range maze {
		for _, val := range row {
			fmt.Print(" ", val, " ")
		}
		fmt.Println()
	}
	path := searchMazeForExit(maze, start, end)
	fmt.Printf("\n\nResult\n\n")
	for _, xy := range path {
		maze[xy[0]][xy[1]] = 8
	}
	for _, row := range maze {
		for _, val := range row {
			fmt.Print(" ", val, " ")
		}
		fmt.Println()
	}
	fmt.Printf("\n%v\n", path)

}
