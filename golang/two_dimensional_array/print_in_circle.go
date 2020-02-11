package main

import "fmt"

func main() {
	var (
		m, n            = 5, 6
		tdarray [][]int = initArray(m, n)
	)

	// i,j对应行列索引
	// num为填充数字
	// circle为环索引
	for i, j, num, circle := 0, 0, 1, 0; num <= m*n; num++ {
		tdarray[i][j] = num
		if i == 0+circle && j == 1+circle {
			// 进入下一个环
			circle++
			i++
			continue
		}
		if 0+circle <= i && i < m-1-circle && j == 0+circle {
			// 左边处理逻辑
			i++
			if i == m-1-circle {
				continue
			}
		}
		if i == m-1-circle && 0+circle <= j && j < n-1-circle {
			// 下边处理逻辑
			j++
			if j == n-1-circle {
				continue
			}
		}
		if i > 0+circle && j == n-1-circle {
			// 右边处理逻辑
			i--
			if i == 0+circle {
				continue
			}
		}
		if i == 0+circle && j > 0+circle {
			// 上边处理逻辑
			j--
		}
	}

	printArray(tdarray)
}

func initArray(m, n int) [][]int {
	s := make([][]int, m, m)
	for i := 0; i < m; i++ {
		s[i] = make([]int, n, n)
	}
	return s
}

func printArray(tds [][]int) {
	for _, s := range tds {
		for _, num := range s {
			fmt.Printf("%3d ", num)
		}
		fmt.Println()
	}
}
