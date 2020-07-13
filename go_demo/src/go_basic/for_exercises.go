package main

import "fmt"

func main() {
	/*
		水仙花数：三位数:[100,999]
			每个位上的数字的立方和，刚好等于该数字本身。那么就叫水仙花数
			比如:153
				1x1x1 + 5x5x5 + 3x3x3 =  153
	*/
	for i := 100; i < 1000; i++ {
		num1 := i / 100
		num2 := i / 10 % 10
		num3 := i % 10
		if num1*num1*num1+num2*num2*num2+num3*num3*num3 == i {
			fmt.Println(i)
		}
	}
}
