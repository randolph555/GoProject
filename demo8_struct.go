package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	example1()
	fmt.Println("-------------------------")
	example2()
}

func example1() {
	var p1 Person
	p1.Name = "Tome"
	p1.Age = 30
	fmt.Println("p1 =", p1)

	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2 =", p2)

	p3 := Person{Name: "Aaron", Age: 32}
	fmt.Println("p3 =", p3)

	//匿名结构体
	p4 := struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 33}
	fmt.Println("p4 =", p4)
}

func example2() {
	var p *Person
	p = new(Person) //分配空间
	p.Name = "字符串"
	p.Age = 20
	fmt.Println(*p)
}
