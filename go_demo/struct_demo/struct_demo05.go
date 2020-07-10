package main

import "fmt"

/*
	面向对象:oop
*/
func main() {
	//1、创建父类的对象
	p1 := Person{name: "张三", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	//2、创建子类对象
	s1 := Student{Person{"李四", 17}, "哈尔滨佛学院"}
	fmt.Println(s1)
}

type Person struct {
	name string
	age  int
}

type Student struct {
	Person        //模拟继承结构
	school string //子类新增属性
}
