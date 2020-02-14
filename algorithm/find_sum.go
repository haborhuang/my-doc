package main

import "fmt"

func main() {
	// 找出数组中的两个元素，使其加和等于指定数字
	array := []int{1, 5, 2, 7, 5}
	sum := 10

	// 算法一：简单暴力 两层循环。时间复杂度n*n

	// 算法二：排序数组后，遍历数组找出解。时间复杂度为排序算法复杂度，快排n*log2(n)

	// 算法三：构建二维数组，其下标表示指定两个位置的和。遍历一次数组填充二维数组。时间复杂度n，空间复杂度n*n

	// 算法四：遍历数组构建map，再次遍历数组，查询map中是否存在解。时间复杂度n，空间复杂度n
	m := make(map[int]int, len(array))
	for _, n := range array {
		m[n]++
	}

	for _, n := range array {
		if sum-n == n && m[n] > 1 {
			fmt.Println(n, "+", n)
			break
		}
		if m[sum-n] > 0 {
			fmt.Println(n, "+", sum-n)
			break
		}
	}
}
