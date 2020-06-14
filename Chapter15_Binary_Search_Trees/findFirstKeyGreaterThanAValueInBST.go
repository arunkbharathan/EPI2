package main

func findFirstKeyGreaterThan(bt *binaryTree, key int) int {
	candidate := 0
	for bt != nil {
		if key < bt.data {
			candidate = bt.data
			bt = bt.left
		} else {
			bt = bt.right
		}
	}
	return candidate
}
