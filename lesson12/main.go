package main

import (
	"awesomeProject1/lesson12/book"
	"fmt"
)

func main() {
	fmt.Println(book.BookCheck("红星照耀中国", ""))
	var Student1 = new(book.Student)
	Student1.Name = "dinglz"
	Student1.Age = 19
	Student1.Gender = "Girl"
	fmt.Println(*Student1)
	Student1.ChangeName("Sam")
}
