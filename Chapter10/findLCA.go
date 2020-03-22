package main

func findLCA(bt1, bt2 *binaryTree) *binaryTree {

	l1, l2 := getDepth(bt1), getDepth(bt2)

	if l1 > l2 {
		diff := l1 - l2
		t := bt1
		for i := 0; i < diff; i++ {
			t = t.parent
		}
		x := bt2
		for x != t {
			x = x.parent
			t = t.parent
		}
		return x
	}

	diff := l2 - l1
	t := bt2
	for i := 0; i < diff; i++ {
		t = t.parent
	}
	x := bt1
	for x != t {
		x = x.parent
		t = t.parent
	}
	return x
}

func getDepth(bt1 *binaryTree) int {
	cnt := 0
	t := bt1
	for t.parent != nil {
		cnt++
		t = t.parent
	}
	return cnt
}
