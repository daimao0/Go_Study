package main

import "fmt"

func main() {
	/*
		参数传递：
			A:值传递：传递的是数据的副本。修改数据，对于原始的数据没有影响的
				值类型的数据，默认都是值传递：基础类型,array,structdemo

			B:引用传递:传递的数据地址
				引用类型的数据，默认都是引用传递:slice,map,chan
	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组的数据", arr1)
	fun1(arr1)
	fmt.Println("函数调用后，数组的数据:", arr1)
}
func fun1(arr1 [4]int) {
	fmt.Println("函数中，数组的数据:", arr1)

	arr1[0] = 100
	fmt.Println("函数中，数组的数据的修改后:", arr1)
}
