package main

import "fmt"

/*
	数组指针：首先是一个指针，一个数组的地址
		*[4]Type

	指针数组：首先是一个数组，存储的数据类型是指针
		[4]*Type
*/
func main() {
	//1、创建一个普通数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//2、创建一个指针，存储该数组的地址——>数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)         //&[1 2 3 4]
	fmt.Printf("%p\n", p1)  //数组arr1的的地址
	fmt.Printf("%p\n", &p1) //指针自己的地址
	fmt.Println(*p1)
	fmt.Println(p1[0])

	//3、根据数组指针，操作数组
	(*p1)[0] = 100
	fmt.Println(arr1)

	p1[0] = 200 //简化写法
	fmt.Println(arr1)

	//4、指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)
	fmt.Println(arr3)
	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)
}
