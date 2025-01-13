package main

import "fmt"

// NumList 切片的定义
func NumList() {
	var ArrayListDefault []int
	//空切片()这两个定义是一样的
	var ArrayListEmpty = []int{}
	fmt.Println(ArrayListDefault)
	fmt.Println(ArrayListEmpty)
}

// UsingNumList 切片的使用（添加元素）
func UsingNumList() {
	a := 100
	var ArrayListDefault []int
	//一次性添加一个值
	ArrayListDefault = append(ArrayListDefault, a)
	fmt.Println(ArrayListDefault)
	//一次性添加多个值
	ArrayListDefault = append(ArrayListDefault, 1, 2, 3)
	fmt.Println(ArrayListDefault)
	//一次性添加一个数组(添加一个已定义的数组)
	var Array = []int{4, 5, 6, 7, 8}
	//...代表要对Array进行解包
	ArrayListDefault = append(ArrayListDefault, Array...)
	fmt.Println(ArrayListDefault)
	//一次性添加一个数组(添加一个现场定义的数组)
	ArrayListDefault = append(ArrayListDefault, []int{9, 10, 11}...)
	fmt.Println(ArrayListDefault)
}

// FuncOfNumList 切片的make,len,cap函数
func FuncOfNumList() {
	//[]int 是指类型是int切片，3代表切片现在有三个元素，5代表切片的总容量为5
	var ArrayListDefault = make([]int, 3, 5)
	fmt.Println("切片的内容", ArrayListDefault)
	fmt.Printf("切片里面的元素个数为%d\n", len(ArrayListDefault))
	fmt.Printf("切片的容量大小为%d\n", cap(ArrayListDefault))
}

// NumListFunc 切片的性质
func NumListFunc() {
	//声明一个数组
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7}
	//切片指针的应用
	//p1从索引1开始截取，到索引4前一个元素截止，左开右闭
	var p1 = arr[1:4]
	//p2截取所有元素
	var p2 = arr[:]
	fmt.Println("arr为", arr, "\np1为", p1, "\np2为", p2)
	//可以通过改变指针对应的值来修改数组
	p2[2] = 100
	fmt.Println("修改后的数组值为", arr)
}
func main() {
	NumList()
	UsingNumList()
	FuncOfNumList()
	NumListFunc()
}
