package main

func evenOddMerge(ll1 *ll) *ll {

	var llodd ll
	t := ll1
	r := &llodd
	preT := ll1
	for i := 0; t != nil && t.next != nil; i++ {
		r.next = t.next
		r = r.next
		t.next = r.next
		r.next = nil
		preT = t
		t = t.next
	}
	if preT.next != nil {
		t.next = llodd.next
		return ll1
	}
	preT.next = llodd.next
	return ll1
}
