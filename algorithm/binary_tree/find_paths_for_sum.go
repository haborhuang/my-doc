package main

import "fmt"

// 二叉树根节点到叶子节点所有路径中，找出节点值和为K的所有路径。
func main() {
	/*
			  2
		  3       4
		1   2
	*/
	root := &btree{
		val: 2,
		left: &btree{
			val: 3,
			left: &btree{
				val: 1,
			},
			right: &btree{
				val: 2,
			},
		},
		right: &btree{
			val: 4,
		},
	}
	k := 6

	fmt.Println(findPaths(root, k))
}

type btree struct {
	val   int
	left  *btree
	right *btree
}

func findPaths(tnode *btree, sum int) [][]int {
	return doFindPaths(tnode, 0, sum, []int{}, [][]int{})
}

func doFindPaths(tnode *btree, level int, sum int, path []int, paths [][]int) [][]int {
	if tnode == nil {
		return paths
	}

	path = addNode(level, path, tnode)
	// 叶子节点
	if tnode.left == nil && tnode.right == nil {
		if sum == tnode.val {
			// 添加路径
			paths = append(paths, append(make([]int, 0, level+1), path[:level+1]...))
		}
		return paths
	}

	paths = doFindPaths(tnode.left, level+1, sum-tnode.val, path, paths)
	paths = doFindPaths(tnode.right, level+1, sum-tnode.val, path, paths)
	return paths
}

func addNode(level int, path []int, tnode *btree) []int {
	if len(path) <= level {
		path = append(path, tnode.val)
	} else {
		path[level] = tnode.val
	}
	return path
}
