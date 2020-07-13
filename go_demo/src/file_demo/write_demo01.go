package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		写出数据：

	*/
	fileName := "F:\\test\\test2.txt"
	//step1：打开文件
	//file, err := os.Open(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	//step3:关闭文件
	defer file.Close()

	//step2:写出数据
	bs := []byte{97, 98, 99, 100}
	//n, err := file.Write(bs)
	n, err := file.Write(bs[:2])
	fmt.Println(n)
	HandlerErr(err)

	//直接写出字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandlerErr(err)

}

func HandlerErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
