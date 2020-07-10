package main

import "fmt"

/*
	高阶函数：
		根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数
	fun1(),fun2()
	将fun1函数作为了fun2这个函数的参数

		fun2函数：就叫做高阶函数
			接收了一个函数作为参数的函数，高阶函数
		fun1函数：回调函数
			作为另一个函数的参数的函数，叫做回调函数
*/
func main() {
	//设计一个函数，用于求两个整数的加减乘除运算
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", oper)

	res2 := oper(1, 2, add)
	fmt.Println(res2)
}

func add(a, b int) int {
	return a + b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun) //打印三个参数
	res := fun(a, b)
	return res
}
