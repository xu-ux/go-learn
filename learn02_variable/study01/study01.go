package main

import "fmt"

/**
【变量的初始化】

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。
每个变量会被初始化成其类型的默认值，

例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil。

	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数

*/

/**
【单个变量的声明与赋值】

	-变量的声明格式 var <变量名称> <变量类型>
	-变量的赋值格式 <变量名称> = <表达式> -变量同时声明赋值 var <变量名称> <变量类型> = <表达式>
	-变量自动推断声明赋值 <变量名称> := <表达式>

```
	//声明
	var a int
	//赋值
	a = 100
	//声明并赋值
	var b int = 100
	//自动推断声明赋值
	c := 100
```

*/
func main() {

	// 变量声明
	var name string // 这个是空串 和java不一样java是null

	// 变量的赋值
	var name2 string = "名称"

	// 简短声明,自动推断声明
	x, y := 100, 200

	age := 100

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(x, y)
	fmt.Println("年龄：", age)

	testMultiVar()
}

/**
全局变量
*/
var (
	a = "hello"
	b = "world"
)

/**
多个变量的声明有赋值
-全局变量的声明可使用 var() 的方法简写
-全局变量的声明不可以省略 var ，但可以使用并行方式
-全局变量都可以使用类型推断
-局部变量不可以使用 var() 的方式简写，只能使用并行方式
*/
func testMultiVar() {
	// 多个变量的声明

	fmt.Println(a, b)
	//并行方式
	var age1, age2 int = 20, 23
	//并行自动推断
	name1, name2 := "小王", "小明"

	fmt.Println(name1, "年龄：", age1, name2, "年龄：", age2)

	// 底层打印函数，实际应用中使用fmt包的打印
	print(name1)
	println(name2)
}
