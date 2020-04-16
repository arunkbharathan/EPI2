package main

import "fmt"

func main() {
	threeSumProblem()
}

func threeSumProblem() {
	arr := []int{11, 2, 5, 7, 3}
	sum := 21
	result := threeSum(arr, sum)
	fmt.Printf("In list %v, %v adds upto %d\n", arr, result, sum)
}
