package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

// https://www.mathsisfun.com/games/towerofhanoi.html
func solveTowerOfHanoi(disksToMove int) {
	towers := map[string]*stack.Stack{
		"A": stack.New(),
		"B": stack.New(),
		"C": stack.New(),
	}
	for i := disksToMove; i > 0; i-- {
		towers["A"].Push(i)
	}
	totalMoves := 0
	hanoiRecurssion(disksToMove, towers, "A", "C", "B", &totalMoves)
	fmt.Println("Total moves: ", totalMoves)
}

func hanoiRecurssion(disksToMove int, towers map[string]*stack.Stack, from, to, using string, totalMoves *int) {
	if disksToMove > 0 {
		hanoiRecurssion(disksToMove-1, towers, from, using, to, totalMoves)
		val := towers[from].Pop().(int)
		towers[to].Push(val)
		fmt.Printf("moving %d from  %s to %s\n", val, from, to)
		*totalMoves++
		hanoiRecurssion(disksToMove-1, towers, using, to, from, totalMoves)
	}

}
