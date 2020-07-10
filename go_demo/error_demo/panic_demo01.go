package main

import "fmt"

func main() {

	funA()
	defer myprint("defer main:3...")
	funB()
	defer myprint("defer main:4...")

	fmt.Println("main..over....")
}

func funA() {
	fmt.Println("我是一个函数funA()......")

}
func funB() {
	fmt.Println("我是函数funB()...")
	defer myprint("defer funB():1....")
	for i := 0; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//让程序中断
			panic("funB函数,恐慌了")
		}
	}
	//当外围函数的代码中发生了运行恐慌，只有其中所有的已经defer的函数全部都执行完毕后，该运行恐慌才会真正被扩展至调用处
	defer myprint("defer funB:2...")
}

func myprint(s string) {
	fmt.Println(s)
}
