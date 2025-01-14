package book

import "fmt"

type Student struct {
	Name   string
	Age    int
	Gender string
}

func (this *Student) ChangeName(Name string) {
	this.Name = Name
	fmt.Println(*this)
}
