package main

import "fmt"

func main() {
	funcRAMPath()
	fmt.Println("----------------------")
	funcPoint()
	fmt.Println("---------输出指针数据结果-------------")
	funcArrayPoint()
	fmt.Println("---------交换数据后的结果-------------")
	funcPointInvoke()
	fmt.Println("---------new()函数创建指针-------------")
	func_new_point()
}

/**查看变量内存地址**/
func funcRAMPath() {
	var a int = 10
	fmt.Printf("变量的地址:%x\n", &a)
}

/*
	声明指针
	var ip *int         指向整型
	var fp *float32     指向浮点型
*/
func funcPoint() {
	//实际变量
	var a int = 20
	//指针变量
	var ip *int
	//指针变量存储地址
	ip = &a

	//a的变量地址
	fmt.Printf("a变量的地址是: %x\n", &a)

	//指针变量的存储地址
	fmt.Printf("ip变量存储的指针地址: %x\n", ip)

	//使用指针访问值
	fmt.Printf("*ip 变量的值%d\n", *ip)
}

const MAX int = 3

/*
	数组指针
*/
func funcArrayPoint() {
	a := []int{10, 100, 200}
	var i int
	//声明数组指针变量
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		//整数地址赋值给指针数组
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

/*函数指针传递**/
func funcPointInvoke() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值: %d\n", a)
	fmt.Printf("交换前 b 的值: %d\n", b)

	/**调用函数交换值**/
	swap(&a, &b)

	fmt.Printf("交换后 a 的值: %d\n", a)
	fmt.Printf("交换后 b 的值: %d\n", b)
}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
func swapEasy(a int, b int) {
	a, b = b, a
}

func swapPointEasy(a *int, b *int) {
	*a, *b = *b, *a
}

//new函数 创建指针
func func_new_point() {
	str := new(string)
	*str = "这是一个指针"
	fmt.Println(*str)
}
