package main

import "strings"

func reconstructBinaryTree(inorder, preorder string) *binaryTree {

	if len(preorder) > 0 {
		root := string(preorder[0])
		inorderLR := strings.Split(inorder, root)
		inorderLeft, inorderRight := inorderLR[0], inorderLR[1]
		preorderLeft := preorder[1 : len(inorderLeft)+1]
		preorderRight := preorder[len(inorderLeft)+1 : len(inorderLeft)+len(inorderRight)+1]

		leftSubtree := reconstructBinaryTree(inorderLeft, preorderLeft)
		rightSubtree := reconstructBinaryTree(inorderRight, preorderRight)
		return &binaryTree{0, root, nil, leftSubtree, rightSubtree}

	}
	return nil

}
