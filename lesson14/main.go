package main

import "fmt"

// Animal 假设为父类
type Animal struct {
	Name   string
	Age    int
	Gender string
}

// Dog 匿名字段实现伪继承
type Dog struct {
	Animal
}
type Cat struct {
	Animal
}

// 接口
type functions interface {
	Show()
}

// Show 接口方法的具体实现
func (dog Dog) Show() {
	fmt.Println("叫", dog.Animal.Name, "的狗", dog.Animal.Age, "岁", "性别为", dog.Animal.Gender)
}
func (cat Cat) Show() {
	fmt.Println("叫", cat.Animal.Name, "的猫", cat.Animal.Age, "岁", "性别为", cat.Animal.Gender)
}
func main() {
	F1 := Dog{Animal{"Dinglz", 11, "man"}}
	F2 := Cat{Animal{"Sam", 11, "man"}}
	//:=创建对象可以不用说对象的类型
	F3 := Dog{}
	//就好像Dog继承了父类
	F3.Name = "Dinglz"
	F3.Gender = "man"
	F3.Age = 18
	F1.Show()
	F2.Show()
	F3.Show()
	//通过接口的对象来使用对应的实现类的方法
	var function functions = Dog{}
	function.Show()
}
