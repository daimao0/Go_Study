package main

import "fmt"

func main() {
	/*
		Go中字符串是一个字节的切片
			可以通过将其内容封装在""中来创建字符串，GO中的字符串是Unicode兼容的，并且是UTF-8编码

		字符串是一些字节的集合

		字符:——>对应编码表中的编码值
		A——>65
		B——>66
		a——>97
		字节:byte——>uint8
			utf8
	*/
	//1、定义字符串
	s1 := "hello 中国" //utf8下一个汉字占三个字节
	s2 := "hello world"
	fmt.Println(s1)
	fmt.Println(s2)

	//2、字符串的长度:返回的是字节个数
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	//3、获取某个字节
	fmt.Printf("%c\n", s2[0])

	//4、字符串的遍历
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c", s2[i])
	}
	fmt.Println("\n------------------")
	for i, v := range s1 {
		fmt.Printf("%d----->%c\n", i, v)
	}
	fmt.Println()

	//5、字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1) //根据一个字节切片，构建字符串
	fmt.Println(s3)

	s4 := "abcdef"
	slice2 := []byte(s4)
	fmt.Println(slice2)

	//6、字符串不能修改
	fmt.Println(s4)
	//s4[2]='B'
}
