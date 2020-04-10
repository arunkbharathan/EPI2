package main

func findKLargestelementsIn(bt *binaryTree, k int, result *[]int) {
	if bt == nil {
		return
	}
	findKLargestelementsIn(bt.right, k, result)
	if len(*result) < k {
		*result = append(*result, bt.data)
	} else {
		return
	}
	findKLargestelementsIn(bt.left, k, result)

}
