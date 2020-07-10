package main

import "fmt"

func main() {
	/*.
	数组的排序：
		让数组中的元素具有一定的顺序
	arr:=[5]int{15,23,8,10,7}
		升序:[7,8,10,15,23]
		降序:[23,15,10,8,7]
	排序算法：
	冒泡排序，插入排序，选择排序，希尔排序，堆排序，快速排序。。。

	*/
	arr := [5]int{15, 23, 8, 10, 7}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Print(arr)
}
