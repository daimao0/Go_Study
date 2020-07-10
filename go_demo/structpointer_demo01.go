package main

import "fmt"

/*
	通过指针：
		new(),不是nil，，空指针
			指向了新配的类型的内存空间，里面存储的零值
*/
func main() {
	//1、结构体是值类型
	p1 := Person{name: "李四", age: 20, sex: "女", address: "杭州市"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)

	p2 := p1 //深拷贝

	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	p2.name = "李小花"
	fmt.Println(p2)
	fmt.Println(p1)

	//2、定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)
	fmt.Println(*pp1)
	//(*pp1).name = "王五"
	pp1.name = "王五" //简写
	fmt.Println(pp1)
	fmt.Println(p1)

	//使用内置函数new()，go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	fmt.Printf("%T\n", pp2)
	fmt.Println(pp2)
	pp2.name = "tom"
	pp2.sex = "man"
	pp2.age = 12
	pp2.address = "beijing"
	fmt.Println(pp2)
}

//1、定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
