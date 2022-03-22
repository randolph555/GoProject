package main

import "fmt"

/*
	声明变量必须要使用，否则go编译不通过
    age_2 declared but not usedcompiler
*/
func main() {

	var name string = "name"
	fmt.Println(name)

	var age_1 int8 = 20
	var age_2 int = 30

	//类型推断 :=  语法
	age_3 := 40
	fmt.Println(age_1, age_2, age_3)

	//多变量声明
	var name_2, age_4 = "name2", 50
	fmt.Println(name_2, age_4)

	//多变量声明，类型推断
	name_3, name_4, age_5, age_6 := "name3", "name4", 60, 70
	fmt.Println(name_3, name_4, age_5, age_6)
}
