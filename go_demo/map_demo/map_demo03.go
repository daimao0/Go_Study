package main

import "fmt"

func main() {
	/*
		map和slice的结合使用：
			1、创建map用于存储人的信息
			2、每个map存储一个人的信息
			3、将这些map存入到slice中
			4、打印遍历输出
	*/
	//1、创建map存储第一个人的信息
	map1 := make(map[string]string)
	map1["name"] = "张三"
	map1["age"] = "30"
	map1["sex"] = "男"
	map1["address"] = "北京市xx路xx号"
	fmt.Println(map1)
	//2、创建第二个人的信息
	map2 := make(map[string]string)
	map2["name"] = "李四"
	map2["age"] = "20"
	map2["sex"] = "女"
	map2["address"] = "上海市xx路xx号"
	fmt.Println(map2)
	//3、第三个人的信息
	map3 := map[string]string{"name": "王五", "age": "25", "sex": "男", "address": "石家庄市xx路xx号"}
	fmt.Println(map3)

	//将map存入到slice中
	s1 := make([]map[string]string, 0, 3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	//遍历切片
	for i, v := range s1 {
		//val:map1,map2,map3
		fmt.Printf("第%d个人的信息是:\n", i+1)
		fmt.Printf("\t姓名:%s\n", v["name"])
		fmt.Printf("\t年龄:%s\n", v["age"])
		fmt.Printf("\t性别:%s\n", v["sex"])
		fmt.Printf("\t地址:%s\n", v["address"])
	}
}
