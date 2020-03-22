package main

type binaryTree struct {
	data  int
	left  *binaryTree
	right *binaryTree
}

var (
	bta *binaryTree = &binaryTree{314, nil, nil}

	btb *binaryTree = &binaryTree{6, nil, nil}
	bti *binaryTree = &binaryTree{6, nil, nil}

	btc *binaryTree = &binaryTree{271, nil, nil}
	btf *binaryTree = &binaryTree{561, nil, nil}
	btj *binaryTree = &binaryTree{2, nil, nil}
	bto *binaryTree = &binaryTree{271, nil, nil}

	btd *binaryTree = &binaryTree{28, nil, nil}
	bte *binaryTree = &binaryTree{0, nil, nil}
	btg *binaryTree = &binaryTree{3, nil, nil}
	btk *binaryTree = &binaryTree{1, nil, nil}
	btp *binaryTree = &binaryTree{28, nil, nil}

	bth *binaryTree = &binaryTree{17, nil, nil}
	btl *binaryTree = &binaryTree{401, nil, nil}
	btn *binaryTree = &binaryTree{257, nil, nil}

	btm *binaryTree = &binaryTree{641, nil, nil}
)

func init() {
	bta.left = btb
	bta.right = bti

	btb.left = btc
	btb.right = btf
	bti.left = btj
	bti.right = bto

	btc.left = btd
	btc.right = bte
	// btf.left = nil
	btf.right = btg
	// btj.left = nil
	btj.right = btk
	// bto.left = nil
	bto.right = btp

	// btd.left = nil
	// btd.right = nil
	// bte.left = nil
	// bte.right = nil
	btg.left = bth
	// btg.right = nil
	btk.left = btl
	btk.right = btn
	// btp.left = nil
	// btp.right = nil

	btl.right = btm
}
