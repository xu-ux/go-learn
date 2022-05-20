package main

/**

包的习惯用法：
包名一般是小写的，使用一个简短且有意义的名称。
包名一般要和所在的目录同名，也可以不同，包名中不能包含- 等特殊符号。
包一般使用域名作为目录名称，这样能保证包名的唯一性，比如 GitHub 项目的包一般会放到GOPATH/src/github.com/userName/projectName 目录下。
包名为 main 的包为应用程序的入口包，编译不包含 main 包的源码文件时不会得到可执行文件。
一个文件夹下的所有源码文件只能属于同一个包，同样属于同一个包的源码文件不能放在多个文件夹下。

*/

import (
	_ "go-learn/learn03_package/test"
)
import "go-learn/learn03_package/test2"

func main() {
	//test.Test()
	test2.Test2(test2.Name)
}
