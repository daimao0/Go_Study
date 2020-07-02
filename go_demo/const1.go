package main

import "fmt"

func main() {
	const PI = 3.14
	const PATH = "http:www.baidu.com"

	fmt.Println(PI)
	fmt.Println(PATH)

	const C1, C2, C3 = "c1", "c2", "c3"
	fmt.Println(C1, C2, C3)

	//一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a = 100
		b
	)
	fmt.Println(a, b)

	//枚举类型：使用常量组为枚举类型，一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
