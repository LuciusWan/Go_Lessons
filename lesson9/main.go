package main

import "fmt"

// PtrTest1 声明一个指针，x:="天津"是把字符串首字母的地址给x，然后str获取x的地址，打印*str就是x的值
func PtrTest1() {
	var str *string
	var p **string
	p = &str
	fmt.Println(p)
	x := "天津"
	str = &x
	fmt.Println(*str)
	fmt.Println(str)
	fmt.Println(p)
}

// PtrTest2 new(string)是在堆里面生成一片空间，然后指针str指向这片空间，然后通过*str来为这片空间赋值
func PtrTest2() {
	var str *string = new(string)
	*str = "北京"
	fmt.Println(*str)
	fmt.Println(str)
}

// SortNumList 切片的排序
func SortNumList(ArrayList []int) {
	for i := 0; i < len(ArrayList)-1; i++ {
		for j := i + 1; j < len(ArrayList); j++ {
			if ArrayList[i] > ArrayList[j] {
				var temp = ArrayList[i]
				ArrayList[i] = ArrayList[j]
				ArrayList[j] = temp
			}
		}
	}
}

// SortArray 数组指针的使用及其排序
func SortArray(ArrayList *[5]int) {
	for i := 0; i < len(ArrayList)-1; i++ {
		for j := i + 1; j < len(ArrayList); j++ {
			if ArrayList[i] > ArrayList[j] {
				var temp = ArrayList[i]
				ArrayList[i] = ArrayList[j]
				ArrayList[j] = temp
			}
		}
	}
}
func main() {
	PtrTest1()
	PtrTest2()
	fmt.Println("--------------------------------------------")
	var Array = []int{1, 5, 3, 2, 7}
	SortNumList(Array)
	fmt.Println(Array)
	var Array2 = [5]int{1, 5, 3, 2, 7}
	SortArray(&Array2)
	fmt.Println(Array2)
}
