package main

import "fmt"

func main() {
	funcMapExample()
}

func funcMap() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	p1[2] = "Tom1"
	fmt.Println("p1 :", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	fmt.Println("p2 :", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	fmt.Println("p3 :", p3)

	p4 := map[int]string{}
	p4[1] = "Tom"
	fmt.Println("p4 :", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	fmt.Println("p5 :", p5)

	p6 := map[int]string{
		1: "Tom",
	}
	fmt.Println("p6 :", p6)
}

func funcMapExample() {
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for key := range countryCapitalMap {
		fmt.Println(key, "首都是:", countryCapitalMap[key])
	}

	/**查看元素在集合中是否存在**/
	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 首都是:", capital)
	} else {
		fmt.Println("American 首都不存在")
	}

}
