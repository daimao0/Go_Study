package main

import "fmt"

/*
	空接口（interface{}）：感觉像是Object类
		不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数据
*/
func main() {
	var a1 A = Cat{"花猫"}
	var a2 A = Person{"张三", 22}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	test(a2)
	test(3.14)
	test2(a3)
	test2(a1)

	//map,key字符串，value任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "张三"
	map1["age"] = 30
	map1["friend"] = Person{"Jerry", 18}
	fmt.Println(map1)

	//切片
}
func test(a A) {
	fmt.Println(a)
}
func test2(a interface{}) {
	fmt.Println("--------》", a)
}

//空接口
type A interface {
}

type Cat struct {
	color string
}

type Person struct {
	name string
	age  int
}
