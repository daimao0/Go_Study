package main

import "fmt"

/**
函数的返回值：
	一个函数的执行结果，返回给函数的调用处。执行结果就叫做函数的返回值
return 语句：
	一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处
	函数返回的结果，必须和函数定义的一致：类型，个数，顺序

	1、将函数的结果返回给调用处
	2、同时结束了该函数的执行
空白标识符：专门用于舍弃数据
*/
func main() {
	res1, res2 := rectangle(5, 3)
	fmt.Println("周长:", res1, "面积:", res2)
	res1, _ = rectangle(5, 3)
	fmt.Println("周长:", res1)
}

//函数，用于求矩形的周长和面积
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}
