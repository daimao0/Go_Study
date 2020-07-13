package main

import "fmt"

func main() {
	/*
		递归函数(recursion):一个函数自己调用自己，就叫递归函数
				递归函数要有一个出口，逐渐的向出口靠近
	*/
	//1、求1-5的和
	sum := getSum(5)
	fmt.Println(sum)

	//2
}

/*
	求1-5的和
getSum(5)
	getSum(4)+5
		getSum(3)+4
			getSum(2)+3
				getSum(1)+2
*/
func getSum(n int) int {
	if n > 1 {
		return getSum(n-1) + n
	}
	return 1
}
