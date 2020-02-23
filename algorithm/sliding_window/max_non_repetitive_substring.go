package main

import "fmt"

func main() {
	// 求最长不重复子串
	// 滑动窗口法：扩展窗口右边，直到发现重复字符。发现重复字符时，窗口整体向右移动。

	s := "abcab"
	exists := [255]bool{}
	winLeft, winRight := 0, 0
	max, posL, posR := 0, 0, 0
	for winRight < len(s) {
		// 字符不存在时
		if !exists[s[winRight]] {
			exists[s[winRight]] = true // 标记字符
			winRight++                 // 扩展右边框
			continue
		}

		// 字符已存在时
		// 更新最大值
		if size := winRight - winLeft; size > max {
			max, posL, posR = size, winLeft, winRight
		}

		// 收缩左边框
		exists[s[winLeft]] = false
		winLeft++
	}
	fmt.Println(max)
	fmt.Println(posL, posR)
	fmt.Println(s[posL:posR])
}
