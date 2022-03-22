package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---------------------------------")
	func_goroutine_sleep()
}

//go goroutine 开启一个协程
func func_goroutine_sleep() {
	fmt.Println("func_goroutine_sleep start")
	go func() {
		fmt.Println("sub goroutine")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("func_goroutine_sleep end")
}
