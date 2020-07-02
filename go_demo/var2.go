package main

import "fmt"

func main() {
	var num int
	num = 100
	fmt.Printf("num的数值是:%d,地址是:%p\n",num,&num)
	num = 200
	fmt.Printf("num的数值是:%d,地址是:%p\n",num,&num)

	var s2 []int
	fmt.Println(s2)
	fmt.Println(s2==nil)
}
