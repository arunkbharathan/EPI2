package main

import sll "github.com/emirpasic/gods/lists/singlylinkedlist"

type qData struct {
	bt *binaryTree
	l  int
	r  int
}

func checkBTSatisfiesBSTProperty2(bt *binaryTree) bool {
	const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	r := 1<<(UintSize-1) - 1
	l := -r - 1
	ll := sll.New()
	ll.Append(&qData{bt, l, r})
	for !ll.Empty() {
		elem, _ := ll.Get(0)
		ll.Remove(0)
		item := elem.(*qData)
		if item.bt == nil {
			continue
		}
		if item.l <= item.bt.data && item.bt.data <= item.r {
			ll.Append(&qData{item.bt.left, item.l, item.bt.data})
			ll.Append(&qData{item.bt.right, item.bt.data, item.r})
		} else {
			return false
		}
	}
	return true
}

func checkBTSatisfiesBSTProperty1(bt *binaryTree) bool {
	const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	r := 1<<(UintSize-1) - 1
	l := -r - 1
	val := checkBSTRecursively(bt, l, r)
	return val
}

func checkBSTRecursively(bt *binaryTree, l, r int) bool {
	if bt == nil {
		return true
	}
	if l <= bt.data && bt.data <= r {
		left := checkBSTRecursively(bt.left, l, bt.data)
		right := checkBSTRecursively(bt.right, bt.data, r)
		return left && right
	}
	return false
}
