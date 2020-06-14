package main

import "fmt"

func main() {
	// checkBSTProperty()
	// findFirstKeyGreaterThanAValueInBST()
	// findFirstKLargestElementsInBST()
	// computeLCAInBST()
	// reconstructBST()
	// findMostVisitedPage()
	buildMinimumHeightBSTFromSortedArray()
}

func buildMinimumHeightBSTFromSortedArray() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buildBSTAVL(arr)
	buildBSTRedBlack(arr)
	result := buildMinimumHeightBST(arr)
	fmt.Println(result)

}

// func findMostVisitedPage() {
// 	kMostVisited := 5
// 	mostVisitedPage(kMostVisited)
// }

// func reconstructBST(){
// 	preorder:=[]int{43,23,37,29,31,41,47,53}
// 	bst:=reconstructBSTFromTraversalData(preorder...)
// 	fmt.Println(bst)
// }

// func computeLCAInBST(){
// 	bt1,bt2:=get2BSTNodes()
// 	bt3:=computeLCA(bt1,bt2)
// 	fmt.Printf("LCA is: %d-%s\n",bt3.data,bt3.str)
// }

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
