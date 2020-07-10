package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println("----------")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println(arr)
	fmt.Println("------------")
	for index, value := range arr {
		fmt.Printf("下标是:%d,数值是:%d\n", index, value)
	}
	sum := 0
	for _, value := range arr {
		fmt.Printf("数值是:%d\n", value)
		sum += value
	}
	fmt.Println(sum)
}
