package main

import "fmt"

func main() {
	var array_1 [5]int
	fmt.Println(array_1)

	var array_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_2)

	array_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_3)

	array_4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array_4)

	//设置下标元素，int默认为0
	array_5 := [5]int{0: 1, 1: 3, 4: 88}
	fmt.Println(array_5)

	//二维数组
	var array_6 = [3][5]int{{}, {}, {}}
	fmt.Println(array_6)

	array_7 := [3][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(array_7)

	array_8 := [...][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(array_8)

}
