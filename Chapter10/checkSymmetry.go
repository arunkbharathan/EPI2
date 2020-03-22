package main

import "fmt"

func isBtSymmetric(bt *binaryTree) bool {

	val := checkIfSymmetric(bt.left, bt.right)
	return val
}

func checkIfSymmetric(btl, btr *binaryTree) bool {
	switch {
	case btl == nil && btr == nil:
		return true
	case btl == nil:
		return false
	case btr == nil:
		return false
	case btl.data != btr.data:
		return false
	case btl.data == btr.data:
		a1 := checkIfSymmetric(btl.right, btr.left)
		a2 := checkIfSymmetric(btl.left, btr.right)
		return a1 && a2
	default:
		fmt.Println("Algo Error")
	}
	fmt.Println("Alg Err")
	return false
}
