package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("--------------------example1-------------------")
	exmaple1()
	fmt.Println("--------------------example2-------------------")
	example2()
	fmt.Println("--------------------example3-------------------")
	example3()
	fmt.Println("--------------------example4-------------------")
	example4()
	fmt.Println("--------------------example5-------------------")
	example5()
	fmt.Println("--------------------example6-------------------")
	example6()
	fmt.Println("--------------------example7-------------------")
	example7()
	fmt.Println("--------------------example7-------------------")
	example8()
}

type myString string

type myInt int

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Free interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Skin interface {
	Color() float64
}

//嵌入式接口
type Material interface {
	Free
	Object
}

type Cube struct {
	side float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Student struct {
	Score float64
}

func (f *Student) Area() float64 {
	return f.Score * 1.5
}
func (f Student) Perimeter() float64 {
	return f.Score * 1.5
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

//空类型
func exmaple1() {
	var s Shape
	fmt.Println("value of s is ", s)
	fmt.Printf("value of s is %T\n", s)
}

//同结构多态
func example2() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area if retange s", s.Area())
	fmt.Println("s == r is ", s == r)
}

//不同结构多态
func example3() {
	var s Shape = Rect{10, 3}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("value of s is %0.2f\n\n", s.Area())

	s = Circle{10}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("valye of s is %v\n", s)
	fmt.Printf("valye of s is %0.2f\n", s.Area())
}

//空接口
func example4() {
	ms := myString("hello world!")
	mi := myInt(120)
	r := Rect{10.5, 2.2}
	explain(ms)
	explain(mi)
	explain(r)
}

//空接口
func explain(i interface{}) {
	fmt.Printf("type = %T, value = %v\n", i, i)
}

//类型断言
func example5() {
	var s Free = Cube{3}
	value, ok := s.(Object)
	fmt.Printf("type = %T value = %v  ok = %v\n", s, value, ok)

	value1, ok1 := s.(Skin)
	fmt.Printf("type = %T value = %v ok = %v\n", s, value1, ok1)

}

//类型推断
func example6() {
	explain2("Hello World!")
	explain2(666)
	explain2(Rect{10, 20})
	explain2(true)
}

func explain2(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("type string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("type int ", i)
	default:
		fmt.Printf("type default %T\n", i)
	}
}

//嵌入式接口
func example7() {
	c := Cube{3}
	var m Material = c
	var f Free = c
	var o Object = c
	fmt.Printf("Material type = %T value = %v Area = %v \n", m, m, m.Area())
	fmt.Printf("Free type = %T,value = %v Area =%v \n", f, f, f.Area())
	fmt.Printf("Object type = %T,value = %v Volume = %v\n", o, o, o.Volume())

}

//指针接收
func example8() {
	student := Student{100}
	var s Shape = &student
	area := s.Area()
	perimeter := s.Perimeter()
	fmt.Printf("area =  %0.2f\n", area)
	fmt.Printf("perimeter =  %0.2f\n", perimeter)
}
