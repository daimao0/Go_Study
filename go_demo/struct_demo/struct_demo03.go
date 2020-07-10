package main

import "fmt"

/*
	结构体嵌套：体格结构体的字段，是另一个结构体类型
*/
func main() {
	b1 := Book{}
	b1.bookName = "西游记"
	b1.Price = 45.8
	s1 := Student{}
	s1.name = "张三"
	s1.age = 18
	s1.book = b1
	fmt.Println(b1)
	fmt.Println(s1)

}

//1、定义一个书的结构体
type Book struct {
	bookName string
	Price    float64
}

//2、定义学生的结构体
type Student struct {
	name string
	age  int
	book Book //book深拷贝
}

type Student2 struct {
	name string
	age  int
	book *Book //book的地址,通过设置指针将book浅拷贝
}
