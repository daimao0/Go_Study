package main

import "fmt"

func main() {
	/*
		goto语句：
	*/

	var a = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a++
			goto LOOP
		}
		fmt.Printf("a的值为: %d\n", a)
		a++
	}
}
