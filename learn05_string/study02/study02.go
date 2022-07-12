package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	cal()
}

/**

【计算字符串长度】



len() 函数的返回值的类型为 int，表示字符串的 ASCII 字符个数或字节长度。
Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节，因此使用 len() 获得四个中文文字对应的 12 个字节。

符合得到4个中文字符的长度的结果
Go 语言中 UTF-8 包提供的 RuneCountInString() 函数，统计 Uncode 字符数量。


总结：
ASCII 字符串长度使用 len() 函数。
Unicode 字符串长度使用 utf8.RuneCountInString() 函数。

*/
func cal() {

	title := "狂人日记"
	fmt.Println(len(title))

	title2 := "This is a giant leap for mankind"
	fmt.Println(len(title2))

	// 计算UTF-8的字符个数
	fmt.Println(utf8.RuneCountInString(title))
	fmt.Println(utf8.RuneCountInString("Hello,世界"))
}
