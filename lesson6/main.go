package main

import (
	"fmt"
)

func ArrayTest01() {
	//没有指定数组的内容
	var arr [5]int
	fmt.Println(arr)
	//数组的长度可以用len(arr)来表示
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
	fmt.Println()
}
func ArrayTest02() {
	//指定数组的内容
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

// ArrayTest03 ArrayList同款用法
func ArrayTest03(a int) {
	var arr []int
	fmt.Printf("请输入%d个数字\n", a)
	for i := 0; i < a; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	fmt.Println(arr)
}

// ArrayTest04 二维数组
func ArrayTest04() {
	var arr [5][5]int
	fmt.Println(arr)
}
func main() {
	ArrayTest01()
	ArrayTest02()
	var a int
	fmt.Println("请输入数据个数")
	fmt.Scan(&a)
	ArrayTest03(a)
	ArrayTest04()
}
