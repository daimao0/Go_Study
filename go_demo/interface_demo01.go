package main

import "fmt"

/*
	接口：interface
		在Go中，接口是一组方法签名
		当某个类型为这个接口中的所有方法提供了方法的实现，它被称为实现接口
		Go语言中，接口和类型的实现关系，是非侵入式

		//其他语言中，要显示的定义
		class Mouse implements USB{}
	1、当需要接口类型的对象时，可以使用任意实现类对象代替
	2、接口对象不能访问实现类中的属性
*/
func main() {
	m1 := Mouse{"罗技"}
	fmt.Println(m1.name)
	f1 := FlashDisk{"金士顿"}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)
}

type USB interface {
	start() //USB设备开始工作
	end()   //USB设备结束工作
}
type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标准备就绪，可以开始工作了")
}
func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，可以安全推出....")
}
func (f FlashDisk) start() {
	fmt.Println(f.name, "准备开始工作，可以进行数据的存储")
}
func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出....")
}
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
