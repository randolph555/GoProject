package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("--------------t1----------")
	fmt.Println(t1())
	fmt.Println("--------------t2----------")
	fmt.Println(t2())
	fmt.Println("--------------t3----------")
	fmt.Println(t3())
	fmt.Println("--------------t4----------")
	fmt.Println(t4())
	fmt.Println("--------------example----------")
	example()
	fmt.Println("--------------exit----------")
	exits()
}

//输出 1
func t1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

//传入参数1 ，返回2
func t2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

//传入参数1 ，返回1
func t3() (b int) {
	a := 1
	defer func() {
		a++
	}()
	return 1
}

//返回 1
func t4() (a int) {
	defer func(a int) {
		a++
	}(a)
	return 1
}

func exits() {
	defer fmt.Println("1")
	fmt.Println("exit")
	//当os.Exit()方法退出程序时，defer不会被执行
	os.Exit(0)
}

func example() {
	func_recover()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
}

func func_recover() {
	defer (func() {
		//Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效
		//recover 的宕机恢复机制就对应其他语言中的 try/catch 机制。
		if err := recover(); err != nil {
			fmt.Println("出错了", err)
		}
	})()

	//func_panic()
	//fmt.Printf("这里应该执行不到！")

	//defer 只对当前协程有效
	go func_panic()

}

func func_panic() {
	panic("error")
}
