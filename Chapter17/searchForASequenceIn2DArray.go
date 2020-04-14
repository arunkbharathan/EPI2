package main

import (
	"github.com/emirpasic/gods/sets/hashset"
)

type previousAttempts struct {
	x      int
	y      int
	offset int
}

func search2DArray(seq []int, grid [][]int) bool {
	// can be done with built in maps too
	// c := map[previousAttempts]bool{}
	previouslyChecked := hashset.New()

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if isPatternSuffixContainedAtXY(grid, seq, previouslyChecked, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func isPatternSuffixContainedAtXY(grid [][]int, pattern []int, prevAttempts *hashset.Set, x, y, offset int) bool {
	if len(pattern) == offset {
		// complete
		return true
	}
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return false
	}
	if prevAttempts.Contains(previousAttempts{x, y, offset}) {
		return false
	}
	if grid[x][y] == pattern[offset] && (isPatternSuffixContainedAtXY(grid, pattern, prevAttempts, x+1, y, offset+1) ||
		isPatternSuffixContainedAtXY(grid, pattern, prevAttempts, x-1, y, offset+1) ||
		isPatternSuffixContainedAtXY(grid, pattern, prevAttempts, x, y+1, offset+1) ||
		isPatternSuffixContainedAtXY(grid, pattern, prevAttempts, x+1, y-1, offset+1)) {
		return true

	}
	prevAttempts.Add(previousAttempts{x, y, offset})
	return false
}
