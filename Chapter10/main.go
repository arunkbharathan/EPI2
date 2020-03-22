package main

import "fmt"

func main() {
	checkIfHeightBalanced()
}

func checkIfHeightBalanced() {
	bta := btBalanced()
	fmt.Println(bta)
	fmt.Println(isHeightBalanced(bta))
	fmt.Println()
	bta = btUnbalanced()
	fmt.Println(bta)
	fmt.Println(isHeightBalanced(bta))
}
