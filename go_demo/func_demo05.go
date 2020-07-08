package main

import "fmt"

func main() {
	/*
		函数的返回值：
			一个函数的执行结果，返回给函数的调用出，执行结果就叫做函数的返回值
		return语句:
			一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处
	*/
	//1、设计一个函数，用于求1-n的和，将结果在主函数打印
	sum := getSum(10)
	fmt.Println(sum)
	sum2 := getSum2()
	fmt.Println(sum2)
}
func getSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func getSum2() (sum int) { //定义函数时，指明要返回的数据是哪一个
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}
