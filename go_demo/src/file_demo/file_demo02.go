package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	/*
		文件操作
		1、路径：
			相对路径：relative
			绝对路径：absolute

			.当前目录
			..上一层目录
	*/
	//1、路径
	fileName1 := "F:\\test\\test.txt"
	fileName2 := "test.txt"
	fmt.Println(filepath.IsAbs(fileName1)) //判断是否时绝对路径
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	fmt.Println("获取父目录", path.Join(fileName1, ".."))

	//2、创建目录
	//err := os.Mkdir("F:\\test\\bb", os.ModePerm)
	//if err!=nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("文件夹创建成功")

	//err := os.MkdirAll("F:\\test\\bb\\cc\\dd\\ee", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println("多层文件夹创建成功")

	//3、创建文件: Create采用模式666，创建一个名为name的文件，如果文件已存在会阶段它（为空文件）
	//file1,err:=os.Create("F:\\test\\bb\\abc.txt")
	//if err!=nil {
	//	fmt.Println("err：",err)
	//	return
	//}
	//fmt.Println(file1)

	//4、打开文件
	//file3,err:=os.Open(fileName1)	//只读
	//if err!=nil {
	//	fmt.Println("err:",err)
	//	return
	//}
	//fmt.Println(file3)

	//5、第一个参数，文件名称
	//	第二参数：文件的打开方式
	//	第三个参数：文件的权限：文件不存在创建文件，需要指定文件
	file4, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm) //只读
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)

	//5、关闭文件
	file4.Close()
}
