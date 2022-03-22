package main

import "fmt"

/*
输出方法
fmt.Print：输出到控制台（仅只是输出）
fmt.Println：输出到控制台并换行
fmt.Printf：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）
fmt.Sprintf：格式化并返回一个字符串，不输出。
*/
func main() {
	fmt.Print("输出到控制台不换行")
	fmt.Println("----")
	fmt.Println("输出到控制台并换行")
	fmt.Printf("name=%s,age=%d\n", "Tom", 30)
	fmt.Printf("name=%s,age=%d height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.257))
}
