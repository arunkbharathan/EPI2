package main

func isHeightBalanced(bt *binaryTree) (bool, int) {
	if bt == nil {
		return true, -1
	}
	isL, hL := isHeightBalanced(bt.left)
	if isL == false {
		return false, hL

	}
	isR, hR := isHeightBalanced(bt.right)
	if isR == false {
		return false, hR
	}
	heightDiff := hL - hR
	isBalanced := false
	if heightDiff >= -1 && heightDiff <= 1 {
		isBalanced = true
	}

	return isBalanced, max(hL, hR) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
