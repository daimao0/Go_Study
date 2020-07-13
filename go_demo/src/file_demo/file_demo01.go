package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		FileInfo：文件信息
			interface
				Name()	文件名
				Size()	文件大小，字节为单位
	*/
	fileInfo, err := os.Stat("F:\\test\\test.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	//1、文件名
	fmt.Println(fileInfo.Name())
	//2、文件大小
	fmt.Println(fileInfo.Size()) //字节
	//3、是否为目录
	fmt.Println(fileInfo.IsDir())
	//4、修改时间
	fmt.Println(fileInfo.ModTime())
	//5、权限
	fmt.Println(fileInfo.Mode())
}
