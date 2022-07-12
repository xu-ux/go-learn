package main

import "strings"

/**

字符串索引比较常用的有如下几种方法：

strings.Index：正向搜索子字符串。
strings.LastIndex：反向搜索子字符串。
搜索的起始位置可以通过切片偏移制作。

*/

func testIndex(s, query string) int {

	// 返回的表示query 从 s字符串开始的 ASCII 码位置
	return strings.Index(s, query)
}

func sliceStr(s string, i int) string {
	// 切片 开始位置 : 结束位置
	slice := s[i:]
	return slice
}
