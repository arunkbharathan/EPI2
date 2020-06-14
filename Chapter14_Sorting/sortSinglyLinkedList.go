package main

func sortList1(sll *ll) *ll {
	t := sll.next
	for t != nil && t.next != nil {
		if t.data < t.next.data {
			t = t.next
		} else {
			target := t.next
			tmpHead := sll
			for tmpHead.next.data < target.data {
				tmpHead = tmpHead.next
			}
			tmp := tmpHead.next
			tmpHead.next = target
			tmp2 := target.next
			target.next = tmp
			t.next = tmp2
			// fmt.Println(&tmp2)
		}
	}
	return sll
}

func sortList2(sll *ll) *ll {
	t := sll.next
	if t == nil || t.next == nil {
		ll1 := ll{}
		tt := &ll1
		tt.next = &ll{t.data, nil}
		return tt
		// return sll
	}

	//find mid
	slow := sll.next
	fast := sll.next
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	ll1 := ll{}
	tt1 := &ll1
	tt1.next = slow.next

	ll2 := ll{}
	tt2 := &ll2
	tt2.next = sll.next
	slow.next = nil
	return mergeSortedLinkedList(sortList2(tt1), sortList2(tt2))

}

func mergeSortedLinkedList(ll1, ll2 *ll) *ll {

	ll3 := &ll{}
	ll1tmp := ll1.next
	ll2tmp := ll2.next
	ll3tmp := ll3

	for {
		if ll1tmp == nil && ll2tmp == nil {
			break
		} else if ll2tmp == nil {
			ll3tmp.next = ll1tmp
			ll1tmp = ll1tmp.next
		} else if ll1tmp == nil {
			ll3tmp.next = ll2tmp
			ll2tmp = ll2tmp.next
		} else {
			if ll1tmp.data < ll2tmp.data {
				ll3tmp.next = ll1tmp
				ll1tmp = ll1tmp.next
			} else {
				ll3tmp.next = ll2tmp
				ll2tmp = ll2tmp.next
			}
		}
		ll3tmp = ll3tmp.next
	}

	return ll3
}
