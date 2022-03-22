package main

import "fmt"

func main() {

	var slice_1 []int //nil 切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_1), cap(slice_1), slice_1)

	var slice_2 = []int{} //空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_2), cap(slice_2), slice_2)

	var slice_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_3), cap(slice_3), slice_3)

	slice_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_4), cap(slice_4), slice_4)

	//分配内存主要有内置函数new和make，用于初始化集合类型的对象
	//make函数 param1=type类型 param2=长度 param3=预留长度，长度是5，最高可以达到8，但是不能低于5
	var slice_5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_5), cap(slice_5), slice_5)

	slice_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice_6), cap(slice_6), slice_6)
}
