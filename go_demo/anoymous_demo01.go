package main

import "fmt"

/*
匿名：没有名字
	匿名函数：没有名字的函数

定义一个匿名函数，可以直接调用，也可以赋给变量调用变量
	匿名函数：
		Go语言是支持函数式编程：
		1、将匿名函数作为另一个函数的参数，回调函数
		2、将匿名函数作为另一个函数的返回值，可以形成闭包结构
*/
func main() {

	fun1()
	fun2 := fun1
	fun2()
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	fun3()
}

func fun1() {
	fmt.Println("我是fun1()函数")
}
