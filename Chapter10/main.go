package main

import "fmt"

func main() {
	// 	checkIfHeightBalanced()
	// 	checkIfBTisSymmetric()
	// 	findLowestCommonAncestor()
	// inOrderTraversalUsingLoop()
	reconstructBinaryTreeFromTraversalData()
}

func reconstructBinaryTreeFromTraversalData() {
	inorder := "FBAEHCDIG"
	preorder := "HBFEACDGI"
	reconBTree := reconstructBinaryTree(inorder, preorder)
	fmt.Println(reconBTree)
	inorder = "DCEBFHGAJLMKNIOP"
	preorder = "ABCDEFGHIJKLMNOP"
	reconBTree = reconstructBinaryTree(inorder, preorder)
	fmt.Println(reconBTree)
}

// func inOrderTraversalUsingLoop() {
// 	bt := btUnbalanced()
// 	fmt.Println(bt)
// 	inOrderLoopTraversal(bt)
// }

// // func findLowestCommonAncestor() {
// // 	bt1, bt2 := btGetTwoNodes()
// // 	fmt.Println(findLCA(bt1, bt2).data)
// // }

// func checkIfBTisSymmetric() {
// 	bt := btSymmetric1()
// 	fmt.Println(bt)
// 	fmt.Println("Symmetric:", isBtSymmetric(bt))

// 	bt = btAsymmetric1()
// 	fmt.Println(bt)
// 	fmt.Println(isBtSymmetric(bt))

// 	bt = btAsymmetric2()
// 	fmt.Println(bt)
// 	fmt.Println(isBtSymmetric(bt))

// 	bt = btUnbalanced()
// 	fmt.Println(bt)
// 	fmt.Println(isBtSymmetric(bt))

// 	bt = btBalanced()
// 	fmt.Println(bt)
// 	fmt.Println(isBtSymmetric(bt))
// }

// func checkIfHeightBalanced() {
// 	bta := btBalanced()
// 	fmt.Println(bta)
// 	fmt.Println(isHeightBalanced(bta))
// 	fmt.Println()
// 	bta = btUnbalanced()
// 	fmt.Println(bta)
// 	fmt.Println(isHeightBalanced(bta))
// }
