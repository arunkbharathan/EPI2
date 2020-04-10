package main

import (
	"container/list"
	"fmt"
	"strings"
)

type binaryTree struct {
	data   int
	str    string
	parent *binaryTree
	left   *binaryTree
	right  *binaryTree
}

func (bt *binaryTree) String() string {
	li := list.New()
	li.PushBack(bt)
	var str = &strings.Builder{}
	retstr := printFromQ(li, str)
	return retstr
}

func printFromQ(li *list.List, str *strings.Builder) string {
	cnt := li.Len()
	if cnt == 0 {
		return str.String()
	}
	for i := 0; i < cnt; i++ {
		val := li.Front()
		bn := val.Value.(*binaryTree)
		li.Remove(val)
		if bn != nil {
			// str.WriteString(fmt.Sprintf("%d ", bn.data))
			str.WriteString(fmt.Sprintf("%d ", bn.data))
			li.PushBack(bn.left)
			li.PushBack(bn.right)
		}
	}
	str.WriteString(fmt.Sprintln())
	retval := printFromQ(li, str)
	return retval
}
func bsTree() *binaryTree {
	var (
		bta *binaryTree = &binaryTree{19, "A", nil, nil, nil}

		btb *binaryTree = &binaryTree{7, "B", nil, nil, nil}
		bti *binaryTree = &binaryTree{43, "I", nil, nil, nil}

		btc *binaryTree = &binaryTree{3, "C", nil, nil, nil}
		btf *binaryTree = &binaryTree{11, "F", nil, nil, nil}
		btj *binaryTree = &binaryTree{23, "J", nil, nil, nil}
		bto *binaryTree = &binaryTree{47, "O", nil, nil, nil}

		btd *binaryTree = &binaryTree{2, "D", nil, nil, nil}
		bte *binaryTree = &binaryTree{5, "E", nil, nil, nil}
		btg *binaryTree = &binaryTree{17, "G", nil, nil, nil}
		btk *binaryTree = &binaryTree{37, "K", nil, nil, nil}
		btp *binaryTree = &binaryTree{53, "P", nil, nil, nil}

		bth *binaryTree = &binaryTree{13, "H", nil, nil, nil}
		btl *binaryTree = &binaryTree{29, "L", nil, nil, nil}
		btn *binaryTree = &binaryTree{41, "N", nil, nil, nil}

		btm *binaryTree = &binaryTree{31, "M", nil, nil, nil}
	)
	bta.parent = nil

	bta.left = btb
	btb.parent = bta
	bta.right = bti
	bti.parent = bta

	btb.left = btc
	btc.parent = btb
	btb.right = btf
	btf.parent = btb
	bti.left = btj
	btj.parent = bti
	bti.right = bto
	bto.parent = bti

	btc.left = btd
	btd.parent = btc
	btc.right = bte
	bte.parent = btc
	btf.right = btg
	btg.parent = btf
	btj.right = btk
	btk.parent = btj
	bto.right = btp
	btp.parent = bto

	btg.left = bth
	bth.parent = btg

	btk.left = btl
	btl.parent = btk
	btk.right = btn
	btn.parent = btk

	btl.right = btm
	btm.parent = btl
	return bta
}

func btUnbalanced() *binaryTree {
	var (
		bta *binaryTree = &binaryTree{314, "A", nil, nil, nil}

		btb *binaryTree = &binaryTree{6, "B", nil, nil, nil}
		bti *binaryTree = &binaryTree{6, "I", nil, nil, nil}

		btc *binaryTree = &binaryTree{271, "C", nil, nil, nil}
		btf *binaryTree = &binaryTree{561, "F", nil, nil, nil}
		btj *binaryTree = &binaryTree{2, "J", nil, nil, nil}
		bto *binaryTree = &binaryTree{271, "O", nil, nil, nil}

		btd *binaryTree = &binaryTree{28, "D", nil, nil, nil}
		bte *binaryTree = &binaryTree{0, "E", nil, nil, nil}
		btg *binaryTree = &binaryTree{3, "G", nil, nil, nil}
		btk *binaryTree = &binaryTree{1, "K", nil, nil, nil}
		btp *binaryTree = &binaryTree{28, "P", nil, nil, nil}

		bth *binaryTree = &binaryTree{17, "H", nil, nil, nil}
		btl *binaryTree = &binaryTree{401, "L", nil, nil, nil}
		btn *binaryTree = &binaryTree{257, "N", nil, nil, nil}

		btm *binaryTree = &binaryTree{641, "M", nil, nil, nil}
	)

	bta.parent = nil

	bta.left = btb
	btb.parent = bta
	bta.right = bti
	bti.parent = bta

	btb.left = btc
	btc.parent = btb
	btb.right = btf
	btf.parent = btb
	bti.left = btj
	btj.parent = bti
	bti.right = bto
	bto.parent = bti

	btc.left = btd
	btd.parent = btc
	btc.right = bte
	bte.parent = btc
	btf.right = btg
	btg.parent = btf
	btj.right = btk
	btk.parent = btj
	bto.right = btp
	btp.parent = bto

	btg.left = bth
	bth.parent = btg

	btk.left = btl
	btl.parent = btk
	btk.right = btn
	btn.parent = btk

	btl.right = btm
	btm.parent = btl
	return bta
}

func btBalanced() *binaryTree {
	var (
		bta *binaryTree = &binaryTree{314, "", nil, nil, nil}

		btb *binaryTree = &binaryTree{6, "", nil, nil, nil}
		bti *binaryTree = &binaryTree{6, "", nil, nil, nil}

		btc *binaryTree = &binaryTree{271, "", nil, nil, nil}
		btf *binaryTree = &binaryTree{561, "", nil, nil, nil}
		btj *binaryTree = &binaryTree{2, "", nil, nil, nil}
		bto *binaryTree = &binaryTree{271, "", nil, nil, nil}

		btd *binaryTree = &binaryTree{28, "", nil, nil, nil}
		bte *binaryTree = &binaryTree{0, "", nil, nil, nil}
		btg *binaryTree = &binaryTree{3, "", nil, nil, nil}
		btk *binaryTree = &binaryTree{1, "", nil, nil, nil}
		btp *binaryTree = &binaryTree{28, "", nil, nil, nil}
	)

	bta.parent = nil

	bta.left = btb
	btb.parent = bta
	bta.right = bti
	bti.parent = bta

	btb.left = btc
	btc.parent = btb
	btb.right = btf
	btf.parent = btb
	bti.left = btj
	btj.parent = bti
	bti.right = bto
	bto.parent = bti

	btc.left = btd
	btd.parent = btc
	btc.right = bte
	bte.parent = btc
	btf.right = btg
	btg.parent = btf
	btj.right = btk
	btk.parent = btj
	bto.right = btp
	btp.parent = bto
	return bta
}

func btSymmetric1() *binaryTree {
	bta := &binaryTree{314, "", nil, nil, nil}

	btb := &binaryTree{6, "", nil, nil, nil}
	bte := &binaryTree{6, "", nil, nil, nil}

	btc := &binaryTree{2, "", nil, nil, nil}
	btf := &binaryTree{2, "", nil, nil, nil}

	btd := &binaryTree{3, "", nil, nil, nil}
	btg := &binaryTree{3, "", nil, nil, nil}

	bta.left = btb
	btb.right = btc
	btc.right = btd

	bta.right = bte
	bte.left = btf
	btf.left = btg
	return bta
}

func btAsymmetric1() *binaryTree {
	bta := &binaryTree{314, "", nil, nil, nil}

	btb := &binaryTree{6, "", nil, nil, nil}
	bte := &binaryTree{6, "", nil, nil, nil}

	btc := &binaryTree{561, "", nil, nil, nil}
	btf := &binaryTree{2, "", nil, nil, nil}

	btd := &binaryTree{3, "", nil, nil, nil}
	btg := &binaryTree{1, "", nil, nil, nil}

	bta.left = btb
	btb.right = btc
	btc.right = btd

	bta.right = bte
	bte.left = btf
	btf.left = btg
	return bta
}

func btAsymmetric2() *binaryTree {
	bta := &binaryTree{314, "", nil, nil, nil}

	btb := &binaryTree{6, "", nil, nil, nil}
	bte := &binaryTree{6, "", nil, nil, nil}

	btc := &binaryTree{561, "", nil, nil, nil}
	btf := &binaryTree{561, "", nil, nil, nil}

	btd := &binaryTree{3, "", nil, nil, nil}

	bta.left = btb
	btb.right = btc
	btc.right = btd

	bta.right = bte
	bte.left = btf
	return bta
}

func btGetTwoNodes() (*binaryTree, *binaryTree) {
	var (
		bta *binaryTree = &binaryTree{314, "", nil, nil, nil}

		btb *binaryTree = &binaryTree{6, "", nil, nil, nil}
		bti *binaryTree = &binaryTree{7, "", nil, nil, nil}

		btc *binaryTree = &binaryTree{271, "", nil, nil, nil}
		btf *binaryTree = &binaryTree{561, "", nil, nil, nil}
		btj *binaryTree = &binaryTree{2, "", nil, nil, nil}
		bto *binaryTree = &binaryTree{272, "", nil, nil, nil}

		btd *binaryTree = &binaryTree{28, "", nil, nil, nil}
		bte *binaryTree = &binaryTree{0, "", nil, nil, nil}
		btg *binaryTree = &binaryTree{3, "", nil, nil, nil}
		btk *binaryTree = &binaryTree{1, "", nil, nil, nil}
		btp *binaryTree = &binaryTree{28, "", nil, nil, nil}

		bth *binaryTree = &binaryTree{17, "", nil, nil, nil}
		btl *binaryTree = &binaryTree{401, "", nil, nil, nil}
		btn *binaryTree = &binaryTree{257, "", nil, nil, nil}

		btm *binaryTree = &binaryTree{641, "", nil, nil, nil}
	)

	bta.parent = nil

	bta.left = btb
	btb.parent = bta
	bta.right = bti
	bti.parent = bta

	btb.left = btc
	btc.parent = btb
	btb.right = btf
	btf.parent = btb
	bti.left = btj
	btj.parent = bti
	bti.right = bto
	bto.parent = bti

	btc.left = btd
	btd.parent = btc
	btc.right = bte
	bte.parent = btc
	btf.right = btg
	btg.parent = btf
	btj.right = btk
	btk.parent = btj
	bto.right = btp
	btp.parent = bto

	btg.left = bth
	bth.parent = btg

	btk.left = btl
	btl.parent = btk
	btk.right = btn
	btn.parent = btk

	btl.right = btm
	btm.parent = btl
	return btm, btb
}

func get2BSTNodes() (*binaryTree, *binaryTree) {
	var (
		bta *binaryTree = &binaryTree{19, "A", nil, nil, nil}

		btb *binaryTree = &binaryTree{7, "B", nil, nil, nil}
		bti *binaryTree = &binaryTree{43, "I", nil, nil, nil}

		btc *binaryTree = &binaryTree{3, "C", nil, nil, nil}
		btf *binaryTree = &binaryTree{11, "F", nil, nil, nil}
		btj *binaryTree = &binaryTree{23, "J", nil, nil, nil}
		bto *binaryTree = &binaryTree{47, "O", nil, nil, nil}

		btd *binaryTree = &binaryTree{2, "D", nil, nil, nil}
		bte *binaryTree = &binaryTree{5, "E", nil, nil, nil}
		btg *binaryTree = &binaryTree{17, "G", nil, nil, nil}
		btk *binaryTree = &binaryTree{37, "K", nil, nil, nil}
		btp *binaryTree = &binaryTree{53, "P", nil, nil, nil}

		bth *binaryTree = &binaryTree{13, "H", nil, nil, nil}
		btl *binaryTree = &binaryTree{29, "L", nil, nil, nil}
		btn *binaryTree = &binaryTree{41, "N", nil, nil, nil}

		btm *binaryTree = &binaryTree{31, "M", nil, nil, nil}
	)
	bta.parent = nil

	bta.left = btb
	btb.parent = bta
	bta.right = bti
	bti.parent = bta

	btb.left = btc
	btc.parent = btb
	btb.right = btf
	btf.parent = btb
	bti.left = btj
	btj.parent = bti
	bti.right = bto
	bto.parent = bti

	btc.left = btd
	btd.parent = btc
	btc.right = bte
	bte.parent = btc
	btf.right = btg
	btg.parent = btf
	btj.right = btk
	btk.parent = btj
	bto.right = btp
	btp.parent = bto

	btg.left = bth
	bth.parent = btg

	btk.left = btl
	btl.parent = btk
	btk.right = btn
	btn.parent = btk

	btl.right = btm
	btm.parent = btl
	return btc, btg
}
