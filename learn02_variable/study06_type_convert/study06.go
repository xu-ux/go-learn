package main

import (
	"fmt"
	"math"
)

/**
【类型转换】

## 如何转换

在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明：
	valueOfTypeB = typeB(valueOfTypeA)
	类型 B 的值 = 类型 B(类型 A 的值)

## 精度丢失

类型转换只能在定义正确的情况下转换成功，
例如从一个取值范围较小的类型转换到一个取值范围较大的类型（将 int16 转换为 int32）。
当从一个取值范围较大的类型转换到取值范围较小的类型时（将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况。

*/
func main() {

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)

	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)

	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))

	/**

	根据输出结果，16 位有符号整型的范围是 -32768～32767，
	而变量 a 的值 1047483647 不在这个范围内。

	1047483647 对应的十六进制为 0x3e6f54ff，转为 int16 类型后，长度缩短一半，也就是在十六进制上砍掉一半，变成 0x54ff，对应的十进制值为 21759。

	浮点数在转换为整型时，会将小数部分去掉，只保留整数部分。

	*/
}
