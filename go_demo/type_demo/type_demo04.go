package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person--->", p.name)
}

//类型别名
type People = Person
type Student struct {
	//嵌入两个结构体
	Person
	People
}

func main() {
	var s Student
	//s.name="张三"	//ambiguous selector s.name
	s.Person.name = "张三"
	s.Person.show()
	fmt.Printf("%T,%T\n", s.Person, s.People)
}
