package main

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
