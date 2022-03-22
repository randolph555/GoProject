package main

import (
	"errors"
	"fmt"
)

func main() {}

func nils() {

	//nil 标识符是不能比较的
	//fmt.Println(nil == nil)

	//nil 不是关键字或保留字
	var nil = errors.New("my god")
	fmt.Println(nil)

	//nil 没有默认类型
	fmt.Printf("%T", nil)
	print(nil)

	//不同类型 nil 的指针是一样的
	var arr []int
	var num *int
	//输出 0x0  0x0
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p", num)

	/*
		不同类型的 nil 是不能比较的
		 var m map[int]string
		 var ptr *int
		 fmt.Println(m == ptr)
	*/

	/*
		两个相同类型的 nil 值也可能无法比较
		var s1 []int
		var s2 []int
		fmt.Println(s1 == s2)
	*/

}
