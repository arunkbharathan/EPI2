package main

type binaryTree struct {
	data   int
	parent *binaryTree
	left   *binaryTree
	right  *binaryTree
}

var (
	bta *binaryTree = &binaryTree{314, nil, nil, nil}

	btb *binaryTree = &binaryTree{6, nil, nil, nil}
	bti *binaryTree = &binaryTree{6, nil, nil, nil}

	btc *binaryTree = &binaryTree{271, nil, nil, nil}
	btf *binaryTree = &binaryTree{561, nil, nil, nil}
	btj *binaryTree = &binaryTree{2, nil, nil, nil}
	bto *binaryTree = &binaryTree{271, nil, nil, nil}

	btd *binaryTree = &binaryTree{28, nil, nil, nil}
	bte *binaryTree = &binaryTree{0, nil, nil, nil}
	btg *binaryTree = &binaryTree{3, nil, nil, nil}
	btk *binaryTree = &binaryTree{1, nil, nil, nil}
	btp *binaryTree = &binaryTree{28, nil, nil, nil}

	bth *binaryTree = &binaryTree{17, nil, nil, nil}
	btl *binaryTree = &binaryTree{401, nil, nil, nil}
	btn *binaryTree = &binaryTree{257, nil, nil, nil}

	btm *binaryTree = &binaryTree{641, nil, nil, nil}
)

func btUnbalanced() *binaryTree {
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
