package main

import (
	"awesomeProject1/function"
	"fmt"
	"reflect"
)

const (
	Str1 = "SHENZHEN"
	Str2 = "BEIJING"
	Str3 = "SHANGHAI"
)
const (
	NANCHANG = iota
	TIANJING
	HANGZHOU
)

func main() {
	fmt.Println(" Hello World")
	fmt.Println(function.Func1(1, 2))
	var a int = 13
	var b int = 14
	fmt.Println(function.Func1(a, b))
	var a1 = 100
	var b1 = 200
	fmt.Println("a1+b1的值为", function.Func1(a1, b1))
	var (
		a2 = "I love you"
		b2 = 3.14159265359
	)
	const a3 = "I hate you"
	fmt.Printf("a3的值为%s\na3的类型是%v\n", a3, reflect.TypeOf(a3))
	fmt.Printf("a2的值为%s\na2的类型是%v\n", a2, reflect.TypeOf(a2))
	fmt.Printf("b2的值为%g\nb2的类型是%v\n", b2, reflect.TypeOf(b2))
	fmt.Println("\n---------------------------------------------------------")
	fmt.Println(Str1)
	fmt.Println(Str2)
	fmt.Println(Str3)
	fmt.Println("\n---------------------------------------------------------")
	fmt.Println(NANCHANG)
	fmt.Println(TIANJING)
	fmt.Println(HANGZHOU)
	fmt.Println("\n---------------------------------------------------------")
	f := 200
	x := 300
	s := "なんだか面白いですね"
	fmt.Println(f)
	fmt.Println(x)
	fmt.Println(s)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
