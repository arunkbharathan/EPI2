package main

func reverseSublistLL(x *ll, i, j int) *ll {
	count := 0
	tmpx := x
	for tmpx != nil && count < i-1 {
		tmpx = tmpx.next
		count++
	}
	beginNode := tmpx
	preNode := beginNode.next
	currNode := preNode.next
	postNode := currNode.next
	count++
	for count < j {
		count++
		currNode.next = preNode

		preNode = currNode
		currNode = postNode
		if postNode != nil {
			postNode = postNode.next
		}

	}
	jNode := preNode
	iNode := beginNode.next
	beginNode.next = jNode
	iNode.next = currNode

	return x
}
