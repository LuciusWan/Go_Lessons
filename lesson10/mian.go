package main

import "fmt"

type Student struct {
	name   string
	age    int
	gender string
}

// GetName Student表示这个方法的归属结构体，GetName方法返回类型是string
func (this Student) GetName() string {
	return this.name
}
func (this Student) GetAge() int {
	return this.age
}
func (this Student) GetGender() string {
	return this.gender
}

// SetGender 要改变对象的内容的时候得传递指针类型，*Student，并且传过来的参数为string类型
func (this *Student) SetGender(gender string) {
	this.gender = gender
}
func (this *Student) SetName(name string) {
	this.name = name
}
func (this *Student) SetAge(age int) {
	this.age = age
}
func main() {
	Student1 := Student{
		name:   "Sam",
		age:    20,
		gender: "male",
	}
	fmt.Println(Student1)
	fmt.Println("-------------------------------")
	Student2 := new(Student)
	Student2.name = "James"
	Student2.age = 20
	fmt.Println(*Student2)
	fmt.Println(Student1.age)
	fmt.Println(Student1.gender)
	fmt.Println(Student1.name)
	var Student3 = new(Student)
	Student3.SetName("dinglz")
	fmt.Println(*Student3)
}
