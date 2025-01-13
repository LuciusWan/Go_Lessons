package main

import (
	"fmt"
	"math"
	"unsafe"
)

// Integer 有符号整型
func Integer() {
	var num8 int8 = math.MaxInt8
	var num16 int16 = math.MaxInt16
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num = math.MaxInt
	fmt.Printf("num8的大小是%d,num8占的内存是%d,num8的类型是%T\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的大小是%d,num16占的内存是%d,num16的类型是%T\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的大小是%d,num32占的内存是%d,num32的类型是%T\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的大小是%d,num64占的内存是%d,num64的类型是%T\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的大小是%d,num占的内存是%d,num的类型是%T\n", num, unsafe.Sizeof(num), num)
}

// ShowFloat 浮点数,%g是科学计数法，%f是正常浮点数
func ShowFloat() {
	var floatnum32 float32 = math.MaxFloat32
	var floatnum64 float64 = math.MaxFloat64
	fmt.Printf("floatnum32的大小是%g,floatnum32占的内存是%d,floatnum32的类型是%T\n", floatnum32, unsafe.Sizeof(floatnum32), floatnum32)
	fmt.Printf("floatnum64的大小是%g,floatnum64占的内存是%d,floatnum64的类型是%T\n", floatnum64, unsafe.Sizeof(floatnum64), floatnum64)

}
func main() {
	Integer()
	ShowFloat()
}
