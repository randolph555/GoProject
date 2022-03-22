package main

import (
	"fmt"
	"time"
)

func main() {
	func_chan()
	fmt.Println("---------------------")
	func_goroutine_chan()
	fmt.Println("---------------------")
	func_goroutine_chan2()
	fmt.Println("---------------------")
	func_goroutine_chan3()
	fmt.Println("---------------------")
	func_goroutine_chan4()
	fmt.Println("---------------------")
	sycn()
}

/*
	声明chan
	声明不带缓冲通道
	ch1 := make(chan string)

	声明带10个缓冲通道
	ch2 := make(chan string, 10)

	声明只读通道
	ch3 := make(<-chan string)

	声明只写通道
	cha4 := make(chan<- string)

	注意：
	不带缓冲的通道，进和出都会阻塞。
	带缓冲的通道，进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。

	关闭 chan
	close(chan)

	close 以后不能再写入，写入会出现 panic
	重复 close 会出现 panic
	只读的 chan 不能 close
	close 以后还可以读取数据
*/
func func_chan() {
	//声明chan
	ch1 := make(chan string, 2)
	//写入chan
	ch1 <- "a"
	ch1 <- "b"

	//读取chan
	val_a, bool := <-ch1
	fmt.Println("val_a =", val_a, " bool =", bool)

	//读取chan
	val_b := <-ch1
	fmt.Println("val_b =", val_b)
}

//1个缓冲通道
func func_goroutine_chan() {
	fmt.Println("main start")
	ch := make(chan string, 1)
	//入队列
	ch <- "a"
	go func() {
		//出队列
		val := <-ch
		fmt.Println("val = ", val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

//没有缓冲通道
func func_goroutine_chan2() {
	fmt.Println("main start")
	ch := make(chan string)
	go func() {
		//入chan
		ch <- "aaa"
	}()
	go func() {
		//出chan
		val := <-ch
		fmt.Println("val = ", val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

//3个缓冲通道,不开启goroutine
//带缓冲的通道，如果长度等于缓冲长度时，再进就会阻塞
func func_goroutine_chan3() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("produce end")
}

//3个缓冲通道 ,开启goroutine
func func_goroutine_chan4() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	go consume(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func consume(ch chan string) {
	for {
		msg := <-ch
		fmt.Println("msg = ", msg)
	}
}

//模拟并发同步阻塞
func sycn() {
	ch := make(chan int)
	//异步执行任务
	go say2("world", ch)
	//主线程执行
	say("hello")
	//阻塞结果
	fmt.Println(<-ch)
}
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string, ch chan int) {
	for i := 0; i < 3; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	//写入chan
	ch <- 0
	close(ch)
}
