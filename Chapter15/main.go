package main

import "fmt"

func main() {
	// checkBSTProperty()
	// findFirstKeyGreaterThanAValueInBST()
	// findFirstKLargestElementsInBST()
	computeLCAInBST()

}
func computeLCAInBST(){
	bt1,bt2:=btGetTwoNodes()
	bt3:=computeLCA(bt1,bt2)
	fmt.Printf("LCA is: %d %s\n",bt3.data,bt3.str)
}

// func findFirstKLargestElementsInBST(){
// 	inp := 3
// 	result:=[]int{}
// 	findKLargestElementsIn(bsTree(), inp,&result)
// 	fmt.Println(inp, "largest :", result)
// }

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
