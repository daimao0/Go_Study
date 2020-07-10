package main

import "fmt"

func main() {

	//1、创建Person类型
	p1 := Person{name: "王二狗", age: 33}
	p1.eat() //父类对象，访问父类的方法
	//2、创建Student类型
	s1 := Student{Person{"Tom", 15}, "哈尔滨佛学院"}
	fmt.Println(s1)
	s1.study()
	s1.eat()
}

//1、定义一个父类
type Person struct {
	name string
	age  int
}

//2、定义一个子类
type Student struct {
	Person
	school string
}

//3、方法
func (p Person) eat() {
	fmt.Println("父类的方法，吃窝窝头....")
}
func (s Student) study() {
	fmt.Println("子类新增的方法，学生学习啦。。。。")
}
func (s Student) eat() {
	fmt.Println("子类重写的方法，吃窝窝头....")
}
