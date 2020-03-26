package main

import "log"

// 计算指定树的节点间最大距离。
func main() {
	root := &tnode{
		val: 5,
		left: &tnode{
			val: 3,
			left: &tnode{
				val: 2,
				left: &tnode{
					val: 1,
				},
			},
			right: &tnode{
				val: 4,
			},
		},
		right: &tnode{
			val: 8,
			left: &tnode{
				val: 6,
				right: &tnode{
					val: 7,
				},
			},
			right: &tnode{
				val: 9,
			},
		},
	}

	log.Println(findMaxDist(root))
}

type tnode struct {
	val   int
	left  *tnode
	right *tnode
}

func findMaxDist(root *tnode) int {
	dist, _ := doFindMax(root)
	return dist
}

// 返回 最大距离 和 最大深度
func doFindMax(root *tnode) (int, int) {
	if root == nil {
		return 0, -1
	}

	// 递归查找左子树和右子树
	distL, depL := doFindMax(root.left)
	distR, depR := doFindMax(root.right)

	// 最大距离：max(子树最大距离的较大值，子树最大深度之和 + 2)
	// 最大深度：max(左子树最大深度, 右子树最大深度) + 1
	return max(distL, distR, depL+depR+2), max(depL, depR) + 1
}

func max(n int, nums ...int) int {
	maxN := n
	for _, n := range nums {
		if maxN < n {
			maxN = n
		}
	}
	return maxN
}
