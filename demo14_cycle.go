package main

import "fmt"

func main() {
	fmt.Println("--------------array---------------")
	array_cycle()
	fmt.Println("--------------slice---------------")
	slice_cycle()
	fmt.Println("---------------map--------------")
	map_cycle()
}

func array_cycle() {
	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)
	fmt.Println("")

	//循环
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}
	fmt.Println("")

	//foreach
	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	//for
	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	//使用空白符
	for _, name := range person {
		fmt.Println("name :", name)
	}
}

func slice_cycle() {
	person := []string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(person), cap(person), person)
	fmt.Println("")

	//循环
	for k, v := range person {
		fmt.Printf("person [%d]: %s\n", k, v)
	}
	fmt.Println("")

	//foreach
	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	//for
	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d] %s\n", i, person[i])
	}
	fmt.Println("")

	//使用空白符
	for _, naem := range person {
		fmt.Println("name :", naem)
	}
}

func map_cycle() {
	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}
	fmt.Printf("len=%d map=%v\n", len(person), person)
	fmt.Println("")

	//循环
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}
	fmt.Println("")

	//foreach
	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	//for
	for i := 1; i <= len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	//使用空白符
	for _, name := range person {
		fmt.Println("name : ", name)
	}
}
