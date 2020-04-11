package main

import (
	"fmt"

	avl "github.com/emirpasic/gods/trees/avltree"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func buildMinimumHeightBST(arr []int) *binaryTree {
	tree := buildBST(arr, 0, len(arr))
	return tree
}

func buildBST(arr []int, l, r int) *binaryTree {

	if l >= r {
		return nil
	}
	mid := l + (r-l)/2
	leftSubtree := buildBST(arr, l, mid)
	rightSubtree := buildBST(arr, mid+1, r)

	root := &binaryTree{arr[mid], "", nil, leftSubtree, rightSubtree}
	return root
}

func buildBSTAVL(arr []int) {

	tree := avl.NewWithIntComparator()
	for _, k := range arr {
		tree.Put(k, "")
	}
	fmt.Println(tree)
}

func buildBSTRedBlack(arr []int) {

	tree := rbt.NewWithIntComparator()
	for _, k := range arr {
		tree.Put(k, "")
	}
	fmt.Println(tree)
}
