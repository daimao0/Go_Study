package main

import "fmt"

func main() {
	/*
		一维数组:存储的多个数组是数值本身
		a1:=[3]int{1,2,3,4}

		二维数组：一维数组中每个值存的是另一个数组
		a2 := [3][4]int {{},{},{},{}}
		该二维数组的长度是3，存储的元素是一维数组，一维数组的元素是数值，每个一维数组长度为4
	*/
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)
	fmt.Printf("二维数组的地址:%p\n", &a2)
	fmt.Printf("二维数组的长度:%d\n", len(a2))
	fmt.Printf("一维数组的地址:%p\n", len(a2[0]))
	fmt.Println(a2[0][3])
	fmt.Println(a2[1][2])
	fmt.Println(a2[2][1])
	//遍历二维数组
	for i := 0; i < len(a2); i++ {
		for j := 0; j < len(a2[i]); j++ {
			fmt.Print(a2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-------")
	// for range遍历二维数组
	for _, arr := range a2 {
		for _, val := range arr {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}

	//打印二维数组地址
	for _, arr := range a2 {
		for _, val := range arr {
			fmt.Printf("二维数组的地址:%p\n", &val)
		}
	}
}
