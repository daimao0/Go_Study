package main

import "fmt"

/*
	数值类型：整形，浮点
		进行运算操作
	字符串：
		可以获取单个字符，截取字串，遍历，strings包下的函数操作
	数组，切片，map
		存储数据，修改数据，获取数据，遍历数据
	函数：

	注意点：
*/
func main() {
	//1、整型
	a := 10
	//运算
	a += 5
	fmt.Println("a:", a)

	fmt.Printf("%T\n", fun1)
	fmt.Println(fun1) //fun1看作函数名对应的函数体地址

	//4、直接定义一个函数类型的变量
	var c func(int, int)
	fmt.Println(c)
	c = fun1 //将fun1的地址赋值给C
	fmt.Println(c)

	fun1(10, 100)
	c(20, 100)

	res1 := fun2
	res2 := fun2(1, 2)
	fmt.Println(res1)
	fmt.Println(res2)
}
func fun2(a, b int) int {
	return a + b
}
func fun1(a, b int) {
	fmt.Printf("a:%d ,b:%d\n", a, b)
}
