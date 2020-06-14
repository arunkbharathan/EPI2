package main

func findKLargestElementsIn(bt *binaryTree, k int, result *[]int) {
	if bt == nil {
		return
	}
	findKLargestElementsIn(bt.right, k, result)
	if len(*result) < k {
		*result = append(*result, bt.data)
	} else {
		return
	}
	findKLargestElementsIn(bt.left, k, result)

}
