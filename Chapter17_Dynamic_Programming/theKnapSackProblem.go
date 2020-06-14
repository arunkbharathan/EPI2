package main

func knapSackMaximumValue(W, P []int, weightConstraint int) (int, []int) {
	//result[len(P)][weightConstraint+1]=-1
	result := make([][]int, len(P))
	selected := make([][]bool, len(P))
	for i := range result {
		result[i] = make([]int, weightConstraint+1)
		selected[i] = make([]bool, weightConstraint+1)
		for j := range result[i] {
			result[i][j] = -1
		}
	}

	addInSack(weightConstraint, len(W)-1, W, P, &result, &selected)

	finalList := make([]int, 0, len(P))
	for i := len(P) - 1; i >= 0; {
		// outer:
		for j := weightConstraint; i >= 0 && j >= 0; {
			if selected[i][j] {
				finalList = append(finalList, i)
				j -= W[i]
				i--
			} else {
				i--
			}
		}
	}

	return result[len(result)-1][len(result[0])-1], finalList
}

func addInSack(capacity, k int, W, P []int, result *[][]int, selected *[][]bool) int {
	array := *result
	selectedOrNot := *selected
	if k < 0 {
		return 0
	}
	if array[k][capacity] == -1 {
		withCurrentItem := 0
		withoutCurrentItem := addInSack(capacity, k-1, W, P, result, selected)
		if capacity-W[k] >= 0 {
			withCurrentItem = addInSack(capacity-W[k], k-1, W, P, result, selected) + P[k]
		} else {
			withCurrentItem = 0
		}
		// array[k][capacity] = max(withCurrentItem, withoutCurrentItem)
		if withCurrentItem > withoutCurrentItem {
			array[k][capacity] = withCurrentItem
			selectedOrNot[k][capacity] = true
		} else {
			array[k][capacity] = withoutCurrentItem
		}
	}

	return array[k][capacity]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
