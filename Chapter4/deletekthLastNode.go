package main

func deleteKthLastNode(ll1 *ll, kth int) *ll {

	t := ll1
	cnt := 0
	for t.next != nil {
		if cnt == kth {
			break
		}
		cnt++
		t = t.next
	}
	if cnt != kth {
		return &ll{}
	}
	m := ll1
	var l *ll
	for t.next != nil {
		t = t.next
		l = m
		m = m.next
	}
	l.next = m.next
	return ll1
}
