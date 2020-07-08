package main

import "fmt"

func main() {
	/*
		可变参数:
			概念：一个函数的参数的类型确定，但是个数不确定，就可以使用可变参数
			语法：
				参数名...参数类型
			注意事项:
				A:如果一个函数的参数是可变参数，同时还有其他的参数。可变参数要放在
				B:一个函数的参数列表中最多只能有一个可变参数
	*/

	getSum(1, 2, 3, 4, 5)

	//切片
	s1 := []int{1, 2, 3, 4, 5}
	getSum(s1...)
}

func getSum(nums ...int) {
	fmt.Printf("%T\n", nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("总和是", sum)
}

//只能有一个可变参数，且要写在最后面
func fun1(s1, s2 string, nums ...float64) {

}
