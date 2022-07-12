package main

import "fmt"

/**

strings.Index：正向搜索子字符串。
strings.LastIndex：反向搜索子字符串。
搜索的起始位置可以通过切片偏移制作。
*/
func main() {
	var title string = "你好，世界"
	i := testIndex(title, "，") // 012你 345好 678,
	fmt.Printf("查找的index: %d \n", i)

	// 切片
	slice := sliceStr(title, i)
	fmt.Printf("查找: %s \n", slice)
}
