package main

import "fmt"

/*
	指针作为参数：
	参数的传递:值传递，引用传递

*/
func main() {
	a := 10
	fmt.Println("fun1()函数调用前，a:", a)
	fun1(a)
	fmt.Println("fun1()函数调用后,a:", a)

	fun2(&a)
	fmt.Println("fun2()函数调用后，a:", a)

}

func fun1(num int) {
	fmt.Println("fun1()函数中的num值：", num)
	num = 100
	fmt.Println("fun1()函数中修改num", num)
}

func fun2(p1 *int) {
	fmt.Println("fun2()函数中,p1:", *p1)
	*p1 = 200
	fmt.Println("fun2()函数中，修改p1:", p1)
}
