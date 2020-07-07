package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
			map的遍历：
				使用：for range
		 			数组,切片:	index，value
					map:	key，value
	*/
	map1 := make(map[int]string)
	map1[1] = "张三"
	map1[2] = "李四"
	map1[3] = "王五"
	map1[4] = "赵六"
	map1[5] = "孙七"
	map1[6] = "周八"

	//1、遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println("---------------")
	/*
		如果map的key不连续，下面的遍历就不合适了
	*/
	for i := 1; i < len(map1); i++ {
		fmt.Println(i, "--->", map1[i])
	}

	/*
		1、获取所有的key，——>切片/数组
		2、进行排序
		3、遍历key,——>map[key]
	*/
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)
	for k := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	//使用sort包下的排序方法
	sort.Ints(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s1 := []string{"Apple", "windows", "Oracle", "JAVA", "易语言", "acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
