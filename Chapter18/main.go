package main

import "fmt"

func main() {
	// threeSumProblem()
	// majorityProblem()
	gasUpProblem()
}

func gasUpProblem() {
	mpg := 20.0
	cityPump := []float64{50, 20, 5, 30, 25, 10, 10}
	cityDistance := []float64{900, 600, 200, 400, 600, 200, 100}
	city := solveGasUpProblem(mpg, cityPump, cityDistance)
	fmt.Printf("The ample city to start journey is from %d\n", city)
}

// func majorityProblem() {
// 	str := "bacaacaadaaaccbcdccdcaccacccckckkcccc"
// 	ans := findMajorityElement(str)
// 	fmt.Printf("Majority element in string \"%s\" is \"%s\"\n", str, ans)
// }

// func threeSumProblem() {
// 	arr := []int{11, 2, 5, 7, 3}
// 	sum := 21
// 	result := threeSum(arr, sum)
// 	fmt.Printf("In list %v, %v adds upto %d\n", arr, result, sum)
// }
