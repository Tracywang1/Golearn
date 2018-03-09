package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println(c1.getArea())
	var Book1 Books
	var Book2 Books
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(&Book1)

	printBook(&Book2)
}
