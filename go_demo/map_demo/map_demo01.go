package main

import "fmt"

func main() {
	/*
		map：映射，是一种专门用于存储键值对的集合，属于引用类型
		存储特点：
			A:存储的是无序键值对
			B:键不能重复，并且和value--对应的
					map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错

		语法结构：
			1、创建map
				var map1 map[key类型]value类型
					nil map，无法直接使用
				var map2 = make(map[key类型]value类型)

				var map3 = map[key类型]value类型{key:value,key:value,key:value...}

		每种数据类型初始化值：
			int: 0
			float: 0.0--->0
			string:""
			array:[0,0,0,,0,0]

			slice: nil		//切片底层是数组，所以在空的时候可以直接使用
			map: mil		//空的map不能直接使用
	*/

	//1、创建map
	var map1 map[int]string         //只有声明，没有初始化
	var map2 = make(map[int]string) //创建
	var map3 = map[string]int{"GO": 98, "python": 87, "java": 76, "html": 90}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil) //只有声明，没有初始化
	fmt.Println(map2 == nil, "——>map2的容量", len(map2))
	fmt.Println(map3 == nil, "——>map2的容量", len(map3))
	//2、nil map
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}

	//3、存储键值对到map中
	map1[1] = "hello"
	map1[2] = "memeda"
	map1[3] = "ruby"
	map1[100] = "go"
	map1[7] = ""
	//4、获取数据，根据key获取对应的value值
	//根据key获取对应的value，如果key存在，获取数值，如果key不存在，获取的是value值类型的零值
	fmt.Println(map1)
	fmt.Println(map1[3]) //根据key值，获取对应的value
	fmt.Println(map1[10])
	fmt.Println(map1[7])

	v1, ok := map1[40] //判断key在map中存不存在
	if ok {
		fmt.Println("对应数值是：", v1)
	} else {
		fmt.Println("操作的key不存在", v1)
	}

	//5、修改数据
	fmt.Println(map1)
	map1[3] = "java"
	fmt.Println(map1)

	//6、删除数据
	delete(map1, 3)
	fmt.Println(map1)
	delete(map1, 30) //删除不存在的键值对无影响
	fmt.Println(map1)

	//7、长度
	fmt.Println(len(map1))
}
