package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		拷贝文件
	*/
	srcFile := "F:\\test\\test2.txt"
	destFile := "ab.txt"
	total, err := CopyFile1(srcFile, destFile)
	fmt.Println(total, err)
}

//该函数：用于io操作实现文件的拷贝，返回值是拷贝的总数量（字节）
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	//读写
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("报错了")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
