package main

import "fmt"

func main() {
	// 题目：求第k小的数据
	s := []int{5, 1, 3, 2, 6, -20, -2}
	k := 3

	// 思路一：排序获取。时间复杂度O(n*log2(n))，空间复杂度为n
	// qsort(s, 0, len(s)-1)
	// fmt.Println(s[k-1])
	// fmt.Println(s)

	// 思路二：维护k长度的有序数组，每次将数插入数组（二分法查找插入位置）。时间复杂度O(n*log2(k))，空间复杂度为k
	sa := newSortedArray(k)
	for _, n := range s {
		sa.add(n)
		// fmt.Printf("SortedArray after %d added: %v\n", n, sa)
	}
	fmt.Println(sa.max())
	fmt.Println(sa)
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

func newSortedArray(leng int) *sortedArray {
	return &sortedArray{
		len:   leng,
		array: make([]int, 0, leng),
	}
}

func (sa *sortedArray) max() int {
	return sa.array[sa.len-1]
}

func (sa *sortedArray) add(el int) {
	if len(sa.array) == 0 {
		sa.array = append(sa.array, el)
		return
	}

	// 获取当前实际长度
	leng := sa.len
	if leng > len(sa.array) {
		leng = len(sa.array)
	}

	// 二分法查找插入位置
	pos := sa.findPos(el, 0, leng-1)

	// sortedArray已填满时
	if leng == sa.len {
		if pos < 0 || pos >= leng {
			// 无效位置，不插入
			return
		}

		sa.array = append(
			sa.array[0:pos+1],         // 保留前pos+1个数
			sa.array[pos:sa.len-1]..., // pos位置的数至倒数第二个数后移一位
		)
		sa.array[pos] = el // 更新当前位置
		return
	}

	// sortedArray未填满时
	if pos < 0 || pos > leng {
		// 无效位置，不插入
		return
	}
	if pos == leng {
		// 插入末尾
		sa.array = append(sa.array, el)
		return
	}

	sa.array = append(
		sa.array[0:pos+1],     // 保留前pos+1个数
		sa.array[pos:leng]..., // pos位置的数至最后一个数后移一位
	)
	sa.array[pos] = el // 更新当前位置
}

func (sa *sortedArray) findPos(el int, begin int, end int) int {
	if end < begin {
		return -1
	}
	if end == begin {
		if el < sa.array[begin] {
			return begin
		}
		return end + 1
	}

	pos := (end - begin) / 2
	if el < sa.array[pos] {
		return sa.findPos(el, begin, pos-1)
	}
	if el == sa.array[pos] {
		return pos + 1
	}
	return sa.findPos(el, pos+1, end)
}
