package main

import "fmt"

func main() {
	// threeSumProblem()
	majorityProblem()
}

func majorityProblem() {
	str := "bacaacaadaaaccbcdccdcaccacccckckkcccc"
	ans := findMajorityElement(str)
	fmt.Printf("Majority element in string \"%s\" is \"%s\"\n", str, ans)
}

// func threeSumProblem() {
// 	arr := []int{11, 2, 5, 7, 3}
// 	sum := 21
// 	result := threeSum(arr, sum)
// 	fmt.Printf("In list %v, %v adds upto %d\n", arr, result, sum)
// }
