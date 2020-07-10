package main

import "fmt"

/*
	匿名结构体和匿名字段：

	匿名结构体：没有名字的结构体
		在创建匿名结构体时，同时创建对象

	匿名字段：一个结构体的字段没有字段名
		默认使用数据类型作为名字，匿名字段不可重复
	匿名函数:
*/
func main() {
	s1 := Student{name: "张三", age: 22}
	fmt.Println(s1.name, s1.age)

	func() {
		fmt.Println("hello world...")
	}()

	s2 := struct {
		name string
		age  int
	}{name: "李四", age: 19}
	fmt.Println(s2.name, s2.age)

	w2 := Worker{"王五", 22}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)
}

type Worker struct {
	string //匿名字段，默认使用数据类型作为名字，匿名字段不可重复
	int
}
type Student struct {
	name string
	age  int
}
