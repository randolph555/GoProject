package main

import (
	"fmt"
	"sync"
)

func main() {
	func_map()
	fmt.Println("-----------------------------")
	sync_map()
}

func func_map() {
	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}
	fmt.Println("data :", person)

	//删除元素
	delete(person, 2)
	fmt.Println("data :", person)

	//修改元素
	person[2] = "Jack"
	person[3] = "Kevin"
	fmt.Println("data :", person)

	//追加元素
	person[4] = "seven"
	fmt.Println("data :", person)
}

func sync_map() {
	var smap sync.Map

	//将键值保存到sync.Map
	smap.Store("string", 100)
	smap.Store("int", 90)
	smap.Store("float", 80)
	smap.Store("bool", 80)

	//从sync.Map中根据键取值
	fmt.Println(smap.Load("string"))

	//根据键删除对应的键值对
	smap.Delete("int")

	//遍历所有sync.Map中的键值对
	smap.Range(func(k, v interface{}) bool {
		fmt.Println("intertate: ", k, v)
		return true
	})
}
