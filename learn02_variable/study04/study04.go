package main

import "fmt"

/**
【变量交换】

编程最简单的算法之一，莫过于变量交换。交换变量的常见算法需要一个中间变量进行变量的临时保存。

*/
func main() {

	var s = intSlice{11, 6, 100, 12}
	fmt.Println(s.less(0, 2))

}

/**
用传统方法编写变量交换代码如下：
*/
func change1() {
	var a int = 100
	var b int = 200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a, b)
}

/**
在计算机刚发明时，内存非常“精贵”。这种变量交换往往是非常奢侈的。于是计算机“大牛”发明了一些算法来避免使用中间变量：
*/
func change2() {
	var a int = 100
	var b int = 200
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
}

/**
到了Go语言时，内存不再是紧缺资源，而且写法可以更简单。使用 Go 的“多重赋值”特性，可以轻松完成变量交换的任务：
*/

func change3() {
	var a int = 100
	var b int = 200
	// 多重赋值 多重赋值时，变量的左值和右值按从左到右的顺序赋值。
	b, a = a, b
	fmt.Println(a, b)
}

// 以下代码可以忽略

// 多重赋值在Go语言的错误处理和函数返回值中会大量地使用。例如使用Go语言进行排序时就需要使用交换，代码如下

// 将 IntSlice 声明为 []int 类型
type intSlice []int

// 为 IntSlice 类型编写一个 Len 方法，提供切片的长度。
func (p intSlice) len() int {
	return len(p)
}

// 根据提供的 i、j 元素索引，获取元素后进行比较，返回比较结果。
// i是否比j小
func (p intSlice) less(i, j int) bool {
	return p[i] < p[j]
}

// 根据提供的 i、j 元素索引，交换两个元素的值。
func (p intSlice) swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
