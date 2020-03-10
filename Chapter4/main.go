package main

import (
	"bytes"
	"fmt"
)

type ll struct {
	data int
	next *ll
}

func (l *ll) String() string {
	h := l.next
	s := ""
	w := bytes.NewBufferString(s)
	if h == nil {
		fmt.Fprintf(w, " ")
		return w.String()
	}
	for h.next != nil {
		fmt.Fprintf(w, "%d ", h.data)
		h = h.next
	}
	fmt.Fprintf(w, "%d ", h.data)
	fmt.Fprintln(w, "")
	return w.String()
}

func main() {
	// mergeList()
	// reverseSublist()
	// findIfCyclic()
	// findOverlapp()
	// findKthFromLast()
	evensFirstOddsLast()
}

func evensFirstOddsLast() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ll1 := ll{}
	t := &ll1
	for _, ele := range arr {
		t.next = &ll{ele, nil}
		t = t.next
	}
	fmt.Println(&ll1)
	fmt.Println(evenOddMerge(&ll1))
}

// func findKthFromLast() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	ll1 := ll{}
// 	t := &ll1
// 	for _, ele := range arr {
// 		t.next = &ll{ele, nil}
// 		t = t.next
// 	}
// 	fmt.Println(&ll1)
// 	fmt.Println(deleteKthLastNode(&ll1, 5))
// }

// func findOverlapp() {
// 	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	var ll1, ll2 ll
// 	t := &ll1
// 	commonNode := &ll{}
// 	for _, ele := range arr1 {
// 		t.next = &ll{ele, nil}
// 		t = t.next
// 		if ele == 6 {
// 			commonNode = t
// 		}
// 	}
// 	fmt.Println(&ll1)

// 	arr2 := []int{11, 12, 13}
// 	t = &ll2
// 	lastNode := &ll{}
// 	for _, ele := range arr2 {
// 		t.next = &ll{ele, nil}
// 		t = t.next
// 	}
// 	lastNode = t
// 	fmt.Println(&ll2)
// 	lastNode.next = commonNode
// 	fmt.Println(&ll2)

// 	fmt.Println(isOverlapping(&ll1, &ll2))
// }

// func findIfCyclic() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	// a := []int{1, 2, 3, 6, 5, 4, 8}
// 	x := ll{}
// 	t := &x
// 	var cycleNode *ll
// 	for _, val := range a {
// 		t.next = &ll{val, nil}
// 		t = t.next
// 		if val == 6 {
// 			cycleNode = t
// 		}
// 	}
// 	t.next = cycleNode
// 	// fmt.Println(&x)
// 	l, m, n := isCyclic(&x)
// 	fmt.Printf("isCyclic: %v, Cyclicity: %d, Cyclic Node: %d\n", l, m, n)
// }

// func reverseSublist() {
// 	a := []int{1, 2, 3, 4, 5, 6, 11, 10, 9, 8, 7, 12}
// 	// a := []int{1, 2, 3, 6, 5, 4}
// 	x := ll{}
// 	t := &x
// 	for _, val := range a {
// 		t.next = &ll{val, nil}
// 		t = t.next
// 	}
// 	fmt.Println(&x)
// 	fmt.Println(reverseSublistLL(&x, 7, 11))
// }

// func mergeList() {
// 	a := []int{2, 5, 7, 12, 13, 14}
// 	b := []int{3, 11, 20}
// 	x := ll{}
// 	y := ll{}
// 	var t = &x
// 	for _, j := range a {
// 		t.next = &ll{data: j}
// 		t = t.next
// 	}
// 	t = &y
// 	for _, j := range b {
// 		t.next = &ll{data: j}
// 		t = t.next
// 	}

// 	fmt.Println(&x)
// 	fmt.Println(&y)
// 	fmt.Println(mergeSortedLinkedList(&x, &y))
// }
