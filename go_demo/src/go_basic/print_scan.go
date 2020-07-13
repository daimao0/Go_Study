package main

import "fmt"

func main() {
	/*
			输入和输出：
			fmt包:输入、输出

			输出：
				Print()	//打印
				Printf()	//格式化输出
				Println()	//打印换行

			格式化打印占位符：
				%v,原样输出
				%T，打印类型
				%t,bool类型
				%s,字符串
				%f，浮点
				%d,10进制的整数
				%b，2进制的整数
				%O,8进制
				%x,16进制
				%c，打印字符
				%p，打印地址

		输入：
			Scanln()
	*/

	a := 100
	b := 3.14
	c := true
	d := "hello world"
	e := "Ruby"
	f := 'A'
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	fmt.Println("----------------")
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点类型")
	fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y	阻塞式的
	fmt.Printf("x的数值:%d,y的数值:%f\n", x, y)

	fmt.Println("请输入一个整数，一个浮点类型")
	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)

}
