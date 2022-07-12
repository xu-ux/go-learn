package main

import "fmt"

/**
【变量的作用域】

根据变量定义位置的不同，可以分为以下三个类型：

函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数
*/
func main() {

	// 声明局部变量
	var a int = 10
	b := 30
	c := a + b
	fmt.Printf("a = %d, b = %d, c = %d \n", a, b, c)

	var e, f int
	e = 100
	f = 200
	//d = f - e - 5
	fmt.Printf("d = %d, e = %d, f = %d \n", d, e, f)
}

// 全局变量
var d int

// 形式参数
func test(i1, i2 int) int {
	return i1 + i2
}
