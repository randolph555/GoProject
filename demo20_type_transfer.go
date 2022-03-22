package main

import "fmt"

func main() {
	int2float32()
}

func int2float32() {
	var sum int32 = 17
	var count int = 15
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean的值为: %f\n", mean)
}
