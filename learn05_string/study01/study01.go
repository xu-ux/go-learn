package main

import "fmt"

/**
【Go语言字符串】

转义字符
\n：换行符
\r：回车符
\t：tab 键
\u 或 \U：Unicode 字符
\\：反斜杠自身

字符串拼接符“+”


多行字符串（字符串字面量）
“`”

*/
func main() {
	stringDemo1()
	stringDemo2()
}

// 转义字符
func stringDemo1() {
	fmt.Println("你好，世界\r\n你好，中国！")
}

// 字符串拼接符“+”
func stringDemo2() {
	str := "世界人民" +
		"大团结万岁"
	fmt.Printf(str)
}
