package main

import "fmt"

func main() {
	// 题目：求第k小的数据
	s := []int{1, 3, 2, 6, -20, -2}
	k := 2

	// 思路一：排序获取。时间复杂度O(n*log2(n))
	qsort(s, 0, len(s)-1)
	fmt.Println(s[k-1])
	fmt.Println(s)

	// 思路二：维护k长度的有序数组，每次将数插入数组（二分法查找插入位置）。时间复杂度O(n*log2(k))
	// TODO
}

// 快排
func qsort(s []int, left, right int) {
	sepVal, sepPos := s[left], left
	i, j := left, right
	// 确定sepPos的位置，保证其左侧小于等于sepVal，右侧大于等于sepVal
	for i <= j {
		// 从右向左遍历，找到第一个小于sepVal的位置
		for ; j >= sepPos && s[j] >= sepVal; j-- {
		}
		// 交换
		if j >= sepPos {
			s[sepPos] = s[j]
			sepPos = j
		}

		// 从左向右遍历，找到第一个大于sepVal的位置
		for ; i <= sepPos && s[i] <= sepVal; i++ {
		}
		// 交换
		if i <= sepPos {
			s[sepPos] = s[i]
			sepPos = i
		}
	}
	// 确定位置后，赋值sepVal
	s[sepPos] = sepVal
	// 递归排序左侧
	if sepPos > 1 {
		qsort(s, 0, sepPos-1)
	}
	// 递归排序右侧
	if sepPos+1 < right {
		qsort(s, sepPos+1, right)
	}
}

type sortedArray struct {
	len   int
	array []int
}
