package main

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

type indexHeight struct {
	height int
	index  int
}

func computeLargestRectangleUnderSkyline(buildingHeights []int) (int, int, int) {
	stack := arraystack.New()
	maxArea := 0
	left, right := 0, 0
	for i, h := range buildingHeights {
		area := 0
		if stack.Size() == 0 {
			stack.Push(indexHeight{h, i})
		} else {
			tmp, _ := stack.Peek()
			peekH := tmp.(indexHeight).height
			// peekI := tmp.(indexHeight).index
			if h >= peekH {
				stack.Push(indexHeight{h, i})
			} else {
				popI := 0
				for h <= peekH {
					temp, _ := stack.Pop()
					popH := temp.(indexHeight).height
					popI = temp.(indexHeight).index
					area = popH * (i - popI)
					if area > maxArea {
						left, right = popI, i
						maxArea = area
					}
					if stack.Size() == 0 {
						break
					}
					tmp, _ = stack.Peek()
					peekH = tmp.(indexHeight).height
				}
				stack.Push(indexHeight{h, popI})
			}

		}
		// fmt.Println(stack)
	}

	return maxArea, left, right
}
