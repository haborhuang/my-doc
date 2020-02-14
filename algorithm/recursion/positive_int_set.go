package main

import "fmt"

func main() {
	// 找出1到N正整数组成的所有集合
	// n = 1时，[] [1]
	// n = 2时，[] [1]
	//         [2] [1, 2]
	// 规律为 f(n) = f(n-1) + f(n-1)结果的每个集合中增加元素n
	n := 4

	// 递归法
	res := findSets(n)
	for _, s := range res {
		fmt.Println(s)
	}

	// 非递归
	res = [][]int{
		[]int{},
	}
	for i := 1; i <= n; i++ {
		for _, s := range res {
			s2 := clone(s)
			res = append(res, append(s2, i))
		}
	}
	for _, s := range res {
		fmt.Println(s)
	}
}

func findSets(n int) [][]int {
	if n == 1 {
		return [][]int{[]int{}, []int{1}}
	}

	res := findSets(n - 1)
	for _, s := range res {
		s2 := clone(s)
		res = append(res, append(s2, n))
	}
	return res
}

func clone(s []int) []int {
	s2 := make([]int, len(s))
	for i := range s {
		s2[i] = s[i]
	}
	return s2
}
