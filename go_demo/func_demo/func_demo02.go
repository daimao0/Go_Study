package main

import "fmt"

func main() {
	getSum(10)
	getAdd(1, 2)
	getAdd(11, 22)
}

func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func getAdd(a, b int) {
	sum := a + b
	fmt.Println(sum)
}
