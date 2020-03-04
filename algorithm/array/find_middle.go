package main

import "fmt"

// 找出数组中的元素，其左侧数均比它小，右侧数均比它大
func main() {
	s := []int{1, 3, 2, 5, 7, 2, 4, 9, 11, 12}
	flags := make([]int, len(s))

	// 从左向右遍历，comp记录最大值
	comp := s[0]
	for i := 1; i < len(s)-1; i++ {
		if comp < s[i] {
			flags[i]++
			comp = s[i]
		}
	}

	// 从右向左遍历，comp记录最小值
	comp = s[len(s)-1]
	for i := len(s) - 2; i > 0; i-- {
		if s[i] < comp {
			flags[i]++
			comp = s[i]
		}
	}

	for i := 1; i < len(s)-1; i++ {
		if flags[i] == 2 {
			fmt.Printf("find number %d at index %d\n", s[i], i)
		}
	}
}
