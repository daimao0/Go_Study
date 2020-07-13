package main

import "fmt"

func main() {
	/*
		深拷贝：将拷贝的是数据本身
			值类型的数据，默认都是深拷贝：array，int，float，string，bool，structdemo
		浅拷贝：拷贝的是数据 地址
			导致多个变量指向同一块内存
			引用类型的数据，默认都是浅拷贝:slice ,map
			因为切片是引用类型的数据，直接拷贝的是地址
	*/

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	//copy()
	s3 := []int{7, 8, 9}
	fmt.Println("s2----->", s2)
	fmt.Println("s3----->", s3)

	copy(s2, s3)
	fmt.Println("s2----->", s2)
	fmt.Println("s3----->", s3)
}
