package main

import "fmt"

func main() {
	/*
	iota：特殊的常量，可以被编译器自动修改的常量
	每当定义一个const，iota的初始值为0
	每当定义一个常量，就会自动累加1
	知道下一个const出现，清零

	 */
	const(
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const (
		d = iota
		e
	)
	fmt.Println(d)
	fmt.Println(e)

	//枚举
	const (
		MALE=iota
		FEMALE
		UNKNOW
	)
	fmt.Println(MALE,FEMALE,UNKNOW)
}