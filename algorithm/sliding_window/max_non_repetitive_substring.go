package main

import "fmt"

func main() {
	// 求最长不重复子串
	// 滑动窗口法：扩展窗口右边，直到发现重复字符。发现重复字符时，窗口整体向右移动。

	s := "abcab"
	exists := [255]bool{}
	winLeft, winRight := 0, 0
	for ; winRight < len(s); winRight++ {
		if !exists[s[winRight]] {
			exists[s[winRight]] = true // 标记已存在字符
			continue
		}

		exists[s[winLeft]] = false
		winLeft++
		exists[s[winRight]] = true
	}
	fmt.Println(winRight - winLeft)
	fmt.Println(winLeft, winRight)
	fmt.Println(s[winLeft:winRight])
}
