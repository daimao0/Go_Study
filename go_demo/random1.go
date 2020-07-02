package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)
	for i:=0;i<10;i++ {
		num2:= rand.Intn(10)
		fmt.Println(num2)
	}
	//随机种子
	rand.Seed(1000)
	num := rand.Intn(10)
	fmt.Println("--->",num)
		//时间戳，通过时间戳来获得不同的时间种子

	t1:=time.Now()	//秒
	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1)

	timeStamp2:= t1.UnixNano()	//纳秒
	fmt.Println(timeStamp2)


	//获得随机数
	//step1：设置种子数，可以设置成时间戳
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		fmt.Println("--->",rand.Intn(100))
	}
	//获取指定范围的随机数
	num3:=rand.Intn(46)+3	//[3,49)
	fmt.Println(num3)
	num4:=rand.Intn(62)+15	//[15,77)
	fmt.Println(num4)
}
