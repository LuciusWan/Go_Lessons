package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Sum 切片解引用，这里的形参args就是字符串集合，返回长度
func Sum(args ...string) int {
	return len(args)
}

// Sum2 传过来是完整的切片，返回值为切片内的元素个数
func Sum2(ArrayList []string) int {
	return len(ArrayList)
}

// InterfaceSum ArrayList是一个切片，加上...可以解引用，传过来就是一堆实体元素
func InterfaceSum(ArrayList ...interface{}) int {
	var sum = 0
	flag1 := true
	flag2 := true
	flag3 := true
	for index, elem := range ArrayList {
		if reflect.TypeOf(elem).Name() == "int" {
			fmt.Printf("第%d个元素的类型是int\n", index+1)
			if flag1 {
				sum++
			}
			flag1 = false
		}
		if reflect.TypeOf(elem).Name() == "float64" {
			fmt.Printf("第%d个元素的类型是float64\n", index+1)
			if flag2 {
				sum++
			}
			flag2 = false
		}
		if reflect.TypeOf(elem).Name() == "string" {
			fmt.Printf("第%d个元素的类型是string\n", index+1)
			if flag3 {
				sum++
			}
			flag3 = false
		}
	}
	return sum
}

// BookCheck 两个返回类型，一个是字符串，一个是异常类型
// 若两个形参类型相等，则在最后写上类型即可
func BookCheck(BookName, BookAuthor string) (string, error) {
	//创建一个新的异常类型
	if BookName == "" {
		return BookName, errors.New("error图书名称不能为空")
	}
	if BookAuthor == "" {
		return BookAuthor, errors.New("error图书作者名称不能为空")
	}
	return "图书名称是" + BookName + "图书的作者是" + BookAuthor, nil
}
func main() {
	var Strings = []string{
		"我是奶龙",
		"我是奶龙",
		"我才是奶龙",
	}
	fmt.Println(Strings)
	fmt.Println(Sum(Strings...))
	fmt.Println(Sum2(Strings))
	var num = InterfaceSum("我是奶龙", 1, 3.1415926, "dinglz")
	fmt.Println("上面元素的类型个数为", num)
	fmt.Println(BookCheck("C++程序设计", "朱鸣华"))
	fmt.Println(BookCheck("", "dinglz"))
	fmt.Println(BookCheck("红星照耀中国", ""))
}
