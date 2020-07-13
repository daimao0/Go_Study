package main

import (
	"fmt"
	"time"
)

func main() {

	//1、获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	//2、获取指定的时间
	t2 := time.Date(2008, 7, 14, 16, 30, 28, 0, time.Local)
	fmt.Println(t2)

	//3、time->string之间转换
	s1 := t1.Format("2006年1月2日 15:04:05") //必须是这个时间
	fmt.Println(s1)
	s2 := t1.Format("2006/01/2") //必须是这个时间
	fmt.Println(s2)

	//4、根据当前时间获取指定的内容
	year, month, day := t1.Date()
	fmt.Println(day, month, year)

	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	year2 := t1.Year()
	fmt.Println(year2)
	fmt.Println(t1.YearDay())

	//5、时间戳，指定的日期,距离1970年1月1日0时0分0秒的时间差
	t4 := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	timeStamp1 := t4.Unix()
	fmt.Println(timeStamp1)
	timeStamp2 := t1.Unix()
	fmt.Println(timeStamp2)

	fmt.Println(t4.UnixNano())
	fmt.Println(t1.UnixNano())

	//6、时间间隔
	fmt.Println(t1)
	fmt.Println(t1.Add(time.Minute))
	fmt.Println(t1.AddDate(1, 0, 0))
}
