package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	fmt.Println("sli[1] ==", sli[1])
	fmt.Println("sli[:] ==", sli[:])
	fmt.Println("sli[1:] ==", sli[1:])
	fmt.Println("sli[:4] ==", sli[:4])

	fmt.Println("sli[0:3] ==", sli[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3]), cap(sli[0:3]), sli[0:3])

	//0:3:5    0-3 ,5=cap好像是这个意思
	fmt.Println("sli[0:3:5] ==", sli[0:3:5])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[0:3:5]), cap(sli[0:3:5]), sli[0:3:5])
}
