package main

// don't use parent to find lca
func computeLCA(bt1, bt2 *binaryTree) *binaryTree {
	root1 := &binaryTree{}
	root2 := &binaryTree{}
	root := &binaryTree{}
	tmp := bt1
	// using parent to find root alone
	for tmp != nil {
		root1 = tmp
		tmp = tmp.parent
	}
	tmp = bt2
	for tmp != nil {
		root2 = tmp
		tmp = tmp.parent
	}
	if root1 != root2 {
		return nil
	}
	root = root1

	tmp = root
	for tmp != nil {
		if bt1.data > tmp.data && bt2.data > tmp.data {
			tmp = tmp.right
		} else if bt1.data < tmp.data && bt2.data < tmp.data {
			tmp = tmp.left
		} else {
			return tmp
		}
	}

	return nil
}
