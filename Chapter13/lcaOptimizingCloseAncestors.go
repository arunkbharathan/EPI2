package main

func findLCAOptimized(bt1, bt2 *binaryTree) *binaryTree {

	lcaMap := map[*binaryTree]bool{}
	t := bt1
	s := bt2
	if t == s {
		return t
	}
	for i := 0; t != nil || s != nil; i++ {
		if i%2 == 0 && t != nil {
			t = t.parent
			if _, ok := lcaMap[t]; ok {
				return t
			}
			lcaMap[t] = true

		}
		if i%2 == 1 && s != nil {
			s = s.parent
			if _, ok := lcaMap[s]; ok {
				return s
			}
			lcaMap[s] = true

		}
	}

	return nil
}
