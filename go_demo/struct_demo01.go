package main

import "fmt"

/*
	结构体: 是由一系列具有相同类型或不同类型的数据构成的数据集合
		结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”
*/
func main() {
	//1、方法一
	var p1 Person
	fmt.Println(p1)
	p1.name = "张三"
	p1.age = 22
	p1.sex = "男"
	p1.address = "北京"
	fmt.Println(p1)
	fmt.Printf("姓名:%s,年龄:%d,性别:%s,地址:%s\n", p1.name, p1.age, p1.sex, p1.address)

	//2、方法二
	p2 := Person{}
	p2.name = "Ruby"
	p2.age = 22
	p2.sex = "男"
	p2.address = "北京"
	fmt.Println(p2)
	fmt.Printf("姓名:%s,年龄:%d,性别:%s,地址:%s\n", p2.name, p2.age, p2.sex, p2.address)

	//3、方法三
	p3 := Person{name: "李四", age: 20, sex: "女", address: "杭州市"}
	fmt.Println(p3)
}

//1、定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
