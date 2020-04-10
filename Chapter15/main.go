package main

import "fmt"

func main() {
	// checkBSTProperty()
	// findFirstKeyGreaterThanAValueInBST()
	findFirstKLargestElementsInBST()

}

func findFirstKLargestElementsInBST(){
	inp := 5
	result:=[]int{}
	findKLargestelementsIn(bsTree(), inp,&result)
	fmt.Println(inp, "largest :", result)
}

// func findFirstKeyGreaterThanAValueInBST() {
// 	inp := 23
// 	val := findFirstKeyGreaterThan(bsTree(), inp)
// 	fmt.Println(inp, ":", val)
// }

// func checkBSTProperty(){
// 	satisfies:=checkBTSatisfiesBSTProperty1(bsTree())
// 	fmt.Println(satisfies)
// 	satisfies=checkBTSatisfiesBSTProperty1(btUnbalanced())
// 	fmt.Println(satisfies)
// 	satisfies=checkBTSatisfiesBSTProperty2(bsTree())
// 	fmt.Println(satisfies)
// 	satisfies=checkBTSatisfiesBSTProperty2(btUnbalanced())
// 	fmt.Println(satisfies)
// }
