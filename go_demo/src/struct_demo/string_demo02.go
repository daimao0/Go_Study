package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		string 包下的关于字符串的函数
	*/
	s1 := "hello world"
	//1、是否包含指定的内容——>bool
	fmt.Println(strings.Contains(s1, "llo"))
	//2、是否包含chars中的任意一个字符
	fmt.Println(strings.ContainsAny(s1, "llo"))
	//3、统计substr出现了几次
	fmt.Println(strings.Count(s1, "ll"))

	//4、是否以指定内容开头/结尾
	s2 := "20200708课堂笔记.txt"
	if strings.HasPrefix(s2, "2020") {
		fmt.Println("19年的文件")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档")

	}
	//5、索引Index
	//查找substr在s中第一次出现的位置，不存在返回-1
	fmt.Println(strings.Index(s1, "ll"))
	//chars中任意一个字符第一次出现的位置，不存在返回-1
	fmt.Println(strings.IndexAny(s1, "abcde"))
	fmt.Println(strings.IndexAny(s1, "abcdeh"))
	//查找最后一次出现的(从后往前找)
	fmt.Println(strings.LastIndex(s1, "l"))

	//6、字符串拼接
	ss1 := []string{"abc", "hello", "ruby"}
	s3 := strings.Join(ss1, "*")
	fmt.Println(s3)
	//切割
	ss2 := strings.Split(s3, "*")
	fmt.Println(ss2)

	//7、Repeat 重复拼接
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)

	//8、replace 替换指定字符,n代表几个（-1为所有）
	s6 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s6)

	//9、大小写
	s7 := "heLLo WOrld**123"
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))

	/*
		截取子串:
		substring(start,end)->substr
		str[start:end]->substr
			包含start，不包含end下标	左闭右开
	*/
	fmt.Println(s1)
	s8 := s1[:5]
	fmt.Println(s8)
}
