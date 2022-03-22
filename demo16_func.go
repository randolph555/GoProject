package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"time"
)

func main() {
	str := "12345"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))

	fmt.Println("-----------------------")
	fmt.Printf("current time str : %s\n", getTimeStr())

	fmt.Println("-----------------------")
	func_actual()

	fmt.Println("-----------------------")
	func_callback()

	fmt.Println("-----------------------")
	useClosure()

	fmt.Println("-----------------------")
	func_method()

	fmt.Println("-----------call------------")
	funcCall()

	fmt.Println("-----------funcSkill------------")
	funcSkill()
}

func func1(input string, input2 string, input3 int) (string, string, int) {
	return input, input2, input3
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func getTimeStr() string {
	return time.Now().Format("2022-01-02 15:04:05")
}

//函数作为实参
func func_actual() {
	//声明函数变量
	getSquarRoot := func(x float64) float64 {
		//9的平方根3
		return math.Sqrt(x)
	}
	//使用函数
	fmt.Println(getSquarRoot(9))
}

/**回调函数**/
type cd func(int) int

func func_callback() {
	funcTestCallback(1, callback)
	funcTestCallback(2, func(x int) int {
		fmt.Printf("我是回调,x:%d\n", x)
		return x
	})
}

func funcTestCallback(x int, f cd) {
	f(x)
}

func callback(x int) int {
	fmt.Printf("我是回调,x:%d\n", x)
	return x
}

/**闭包**/
//(1)return返回函数，(2)return返回函数类型
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**调用闭包函数**/
func useClosure() {
	nextNumber := getSequence()
	//调用nextNumber函数 i 变量自增1并返回
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	//创建新的函数nextNumber1,查看结果
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2
	fmt.Println(nextNumber1()) //3
}

type Circle struct {
	radius float64
}

/**使用函数方法**/
func func_method() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

/**函数方法**/
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

//带返回参数名的函数
func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

func funcs() {
	var f func()
	f = fire
	f()
}

//函数变量
func fire() {
	fmt.Println("fire")
}

//匿名函数用作回调函数
func funcCall() {
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

func visit(list []int, f func(v int)) {
	for _, v := range list {
		f(v)
	}
}

var skillParam = flag.String("skill", "", "skill to perform")

//匿名函数实现操作封装
//执行 go run demo16_func.go --skill=fire
func funcSkill() {
	//解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("hello fire")
		},
		"run": func() {
			fmt.Println("hello run")
		},
		"fly": func() {
			fmt.Println("hello fly")
		},
	}
	//skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}