package main

// BST can't be reconstructed from inorder data alone
// but possible from any one of preorder or postorder
// below solution is with preorder
func reconstructBSTFromTraversalData(ints ...int) *binaryTree {

	if len(ints) == 0 {
		return nil
	}
	if len(ints) == 1 {
		return &binaryTree{ints[0], "", nil, nil, nil}
	}
	if len(ints) == 2 && ints[0] > ints[1] {
		left := &binaryTree{ints[1], "", nil, nil, nil}
		root := &binaryTree{ints[0], "", nil, left, nil}
		return root
	}
	if len(ints) == 2 && ints[0] < ints[1] {
		right := &binaryTree{ints[1], "", nil, nil, nil}
		root := &binaryTree{ints[0], "", nil, nil, right}
		return root
	}

	rootData := ints[0]
	root := &binaryTree{rootData, "", nil, nil, nil}
	leftSubtree := []int{}
	rightSubtree := []int{}
	lr := ints[1:]
	for _, val := range lr {
		if val < rootData {
			leftSubtree = append(leftSubtree, val)
		} else {
			rightSubtree = append(rightSubtree, val)
		}
	}
	root.left = reconstructBSTFromTraversalData(leftSubtree...)
	root.right = reconstructBSTFromTraversalData(rightSubtree...)

	return root
}
