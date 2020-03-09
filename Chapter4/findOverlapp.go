package main

func isOverlapping(ll1 *ll, ll2 *ll) (bool, int) {
	ll1len := getlenLL(ll1)
	ll2len := getlenLL(ll2)
	if ll1len > ll2len {
		diff := ll1len - ll2len
		t1 := ll1
		for i := diff; i > 0; i-- {
			t1 = t1.next
		}
		t2 := ll2
		for t2.next != nil {
			if t1 == t2 {
				return true, t1.data
			}
			t1 = t1.next
			t2 = t2.next
		}
	} else {
		diff := ll2len - ll1len
		t1 := ll2
		for i := diff; i > 0; i-- {
			t1 = t1.next
		}
		t2 := ll1
		for t2.next != nil {
			if t1 == t2 {
				return true, t1.data
			}
			t1 = t1.next
			t2 = t2.next
		}
	}
	return false, 0
}

func getlenLL(ll1 *ll) int {
	t := ll1
	cnt := 0
	for ; t.next != nil; cnt++ {
		t = t.next
	}
	return cnt
}
