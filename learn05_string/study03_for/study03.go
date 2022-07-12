package main

import "fmt"

/**

【获取每一个字符串元素】

ASCII 字符串遍历直接使用下标。
Unicode 字符串遍历用 for range。

*/

var title = "Hello,世界！"

func main() {

	// 遍历字符串有下面两种写法。
	demo1()
	fmt.Println("=================================")
	demo2()
}

// 遍历每一个ASCII字符
func demo1() {
	for i := 0; i < len(title); i++ {
		fmt.Printf("ascii: %c  %d\n", title[i], title[i])
		// 这种模式下取到的汉字“惨不忍睹”。由于没有使用 Unicode，汉字被显示为乱码。
	}
}

// 按Unicode字符遍历字符串
func demo2() {
	for _, s := range title {
		fmt.Printf("unicode: %c  %d\n", s, s)
	}
}
