package main

import "fmt"

func main() {

	a := 10 //int
	fmt.Printf("%T\n", a)
	b := [4]int{1, 2, 3, 4} //array
	fmt.Printf("%T\n", b)
	c := []int{1, 2, 3, 4} //slice
	fmt.Printf("%T\n", c)
	d := make(map[int]string) //map
	fmt.Printf("%T\n", d)

	fmt.Printf("%T\n", fun1)
	fmt.Printf("%T\n", fun2)
}

func fun1() {

}

func fun2(n int) int {
	return 0
}
