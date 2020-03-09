package main

func isCyclic(x *ll) (bool, int, int) {

	itr1 := x
	itr2 := x
	looped := false

	for itr2.next != nil && itr2.next.next != nil {
		itr1 = itr1.next
		itr2 = itr2.next.next
		if itr1 == itr2 {
			looped = true
			break
		}
	}
	cyclicity := 0
	for {
		itr1 = itr1.next
		itr2 = itr2.next.next
		cyclicity++
		if itr1 == itr2 {
			break
		}
	}
	itr1 = x
	itr2 = x
	for i := cyclicity; i > 0; i-- {
		itr2 = itr2.next
	}
	for {
		itr1 = itr1.next
		itr2 = itr2.next
		if itr1 == itr2 {
			break
		}
	}
	if !looped {
		return false, 0, 0
	}
	return looped, cyclicity, itr1.data
}
