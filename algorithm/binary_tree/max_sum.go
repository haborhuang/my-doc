package main

func main() {
	// 在根节点到叶子节点所有路径中，找出路径上节点值和的最大值
	bt := initTree() // TODO
	findMax(bt, 0)
}

type btree struct {
	val   int
	left  *btree
	right *btree
}

func findMax(tn *btree, acc int) int {
	if tn == nil {
		return acc
	}
	l := findMax(tn.left, acc+tn.val)
	r := findMax(tn.right, acc+tn.val)
	if l > r {
		return l
	}
	return r
}
