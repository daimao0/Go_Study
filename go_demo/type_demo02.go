package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i1 myint
	var i2 = 100
	i1 = 200
	fmt.Println(i1, i2)

	var name mystr
	name = "王二狗"
	var str1 string
	str1 = "李小虎"
	fmt.Println(name, str1)

	//i1 = i2  cannot use i2 (type int) as type myint in assignment
	fmt.Printf("%T,%T,%T,%T", i1, i2, name, str1)

	fmt.Println("--------------------------")
	res1 := fun1()
	fmt.Println(res1(10, 20))

	fmt.Println("-----------------")
	var i3 myint2
	i3 = 1000
	fmt.Println(i3)
	i3 = i2
	fmt.Println(i3)
	fmt.Printf("%T,%T,%T\n", i1, i2, i3)
}

//1、定义一个新的类型
type myint int
type mystr string

//2、定义函数类型
type myfun func(int, int) string

func fun1() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

//3、类型别名
type myint2 = int //不是重新定义的数据类型，只是给int起别名,和int通用
