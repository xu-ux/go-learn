package main

import "fmt"

/**
【变量类型】

布尔型 bool
	-长度 1字节
	-取值范围 true false
	-主要事项 不可使用0/1代替true/false

整型 int/uint
	-根据运行平台肯为32或64位

	8位整型 int8/uint8
		-长度 1字节
		-取值范围 -128~127 0

	字节型 byte( uint8别名 )

	16位整型 int16/uint16
		-长度 2字节
		-取值范围 -32758~32767 0~65535

	32位整型 int16/uint16
		-长度 4字节
		-取值范围 -2^32/2~2^32/2-1 0~2^32-1

	64位整型 int16/uint16
		-长度 8字节
		-取值范围 -2^64/2~2^64/2-1 0~2^64-1

浮点型 float32/float64
	-长度 4/8字节
	-小数位 精确的7/15小数位

复数 complex64/complex128

足够保存指针的32位或64位整数型 uintptr

其他值类型
	-array, struct, string

引用类型
	-slice, map, chan

接口类型 interface

函数类型 func

*/
func main() {
	valueType()

}

func valueType() {

	/**
	值类型：
	    bool
	    int(32 or 64), int8, int16, int32, int64
	    uint(32 or 64), uint8(byte), uint16, uint32, uint64
	    float32, float64
	    string
	    complex64, complex128
	    array    -- 固定长度的数组
	*/
	// 布尔类型
	var b bool = true
	fmt.Printf("bool：%t \n", b) // %t格式化布尔类型

	// 整数类型 int(32 or 64), int8, int16, int32, int64
	var i1 int8 = -128                       // (-128 ~ 127)
	fmt.Printf("int8：%b int8：%d \n", i1, i1) // %b打印二进制  %d打印10进制

	var i2 int16 = 1550
	fmt.Printf("int16：%b int16：%d \n", i2, i2)

	// 无符号类型
	var ui3 uint8 = 255                          // (0 ~ 255)
	fmt.Printf("uint8：%b uint8：%d \n", ui3, ui3) // %b打印二进制  %d打印10进制
}
