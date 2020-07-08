package main

import "fmt"

func main() { //程序的入口，是一个特殊的函数
	/*
		函数：function
		一、概念：
			具有特定功能的代码，可以被多次调用执行
		二、意义：
			1、可以避免重复的代码
			2、增强程序的扩展性
		三、使用：
			step1、函数的定义
			step2、函数的调用，就是执行函数中代码的过程
		四、语法：
			func funcName(para type1,para type2) (output1 type1 output2 type2){
				return val1,val2;
			}
	*/

	getSum()

}

//定义一个函数：用于求1-10
func getSum() {
	//求1-10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1-10的和是", sum)
}
