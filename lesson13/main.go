package main

import (
	"errors"
	"fmt"
)

func SwitchTest(Grade int) {
	switch {
	case Grade > 90:
		fmt.Println("A")
	case Grade > 80:
		fmt.Println("B")
	case Grade > 70:
		fmt.Println("C")
	case Grade > 60:
		fmt.Println("D")
	default:
		fmt.Println("不及格")
	}
}

// SwitchTest2 Switch的多条件，如果匹配的话直接break无需多余break语句
func SwitchTest2(month int) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31天")
	case 4, 6, 9, 11:
		fmt.Println("30天")
	case 2:
		fmt.Println("28天或29天")
	default:
		fmt.Println("非法输入")
	}
}

// NaiLong case里面的条件可以任意
func NaiLong(str string) (string, error) {
	switch {
	case str == "我是奶龙（1）":
		return "第一句", nil
	case str == "我是奶龙（2）":
		return "第二句", nil
	case str == "我才是奶龙":
		return "第三句", nil
	default:
		return "", errors.New("这不是奶龙视频里的句子")
	}

}
func LoopTest(str string) {
	//range的for循环
	for index, value := range str {
		fmt.Printf("index为%d, value为   %c\n", index, value)
	}
	//不管索引值的range循环
	for _, value := range str {
		fmt.Printf("value为   %c\n", value)
	}
	//正常循环
	for i := 0; i < 5; i++ {
		fmt.Println("我是奶龙")
	}
	//死循环
	for {
		fmt.Println("我才是奶龙")
		break
	}
	//相当于while循环
	i := 0
	for i < 5 {
		fmt.Println("今夜星光闪闪")
		i++
	}
}
func main() {
	//defer关键字可以让语句在main函数执行完之后执行
	//所有的defer语句处理好后入栈，到main函数结束后统一执行
	//defer在释放资源的时候也非常好用
	var str = "我是奶龙250"
	defer fmt.Println("我是奶龙1")
	defer fmt.Println("我是奶龙2")
	defer fmt.Println("我是奶龙3")
	//此语句最后输出结果还是"我是奶龙250"
	defer fmt.Println(str)
	str = "我才是奶龙250"
	var a = 90
	SwitchTest(a)
	var month = 2
	SwitchTest2(month)
	var str1, err1 = NaiLong("我才是奶龙")
	var str2, err2 = NaiLong("我不是奶龙")
	fmt.Println(str1, err1)
	fmt.Println(str2, err2)
	LoopTest("今夜星光闪闪，我爱你的心满满")
}
