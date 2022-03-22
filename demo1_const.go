package main

import "fmt"

/*
	声明常量 const关键字,创建后不可修改
*/
func main() {
	const name string = "name"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2 = "name1", "name2"
	fmt.Println(name_1, name_2)

	const name3, age1 = "name3", 30
	fmt.Println(name3, age1)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(WednesDay)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}

type WeekDay int

const (
	Sunday    WeekDay = iota //iota    每一个有常量声明的行加一   //0
	Monday                   //1
	Tuesday                  //2
	WednesDay                //3
	Thursday                 //4
	Friday                   //5
	Saturday                 //6
)
