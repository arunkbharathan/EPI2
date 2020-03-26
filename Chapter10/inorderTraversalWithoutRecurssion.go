package main

import "fmt"

func inOrderLoopTraversal(bt *binaryTree) {
	currentNode := bt
	var preVNode *binaryTree
	for currentNode != nil {
		if preVNode == nil {
			if currentNode.left == nil {
				fmt.Println(currentNode.str)
				if currentNode.right == nil {
					preVNode = currentNode
					currentNode = currentNode.parent
					continue
				} else {
					preVNode = currentNode
					currentNode = currentNode.right
					continue
				}
			} else {
				preVNode = currentNode
				currentNode = currentNode.left
				continue
			}
		}

		if preVNode == currentNode.parent {
			if currentNode.left == nil {
				fmt.Println(currentNode.str)
				if currentNode.right == nil {
					preVNode = currentNode
					currentNode = currentNode.parent
					continue
				} else {
					preVNode = currentNode
					currentNode = currentNode.right
					continue
				}
			} else {
				preVNode = currentNode
				currentNode = currentNode.left
				continue
			}
		}
		if preVNode == currentNode.left {
			fmt.Println(currentNode.str)
			if currentNode.right == nil {
				preVNode = currentNode
				currentNode = currentNode.parent
				continue
			} else {
				preVNode = currentNode
				currentNode = currentNode.right
				continue
			}
		}
		if preVNode == currentNode.right {
			preVNode = currentNode
			currentNode = currentNode.parent
			continue
		}

	}
}
