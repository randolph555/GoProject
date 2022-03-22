package main

import "fmt"

func main() {
	fmt.Println("---------break---------")
	func_break()
	fmt.Println("---------continue---------")
	func_continue()
	fmt.Println("---------goto---------")
	func_goto()
	fmt.Println("---------switch---------")
	func_switch()

}

func func_break() {
	for i := 1; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("i =", i)
	}
}

func func_continue() {
	for i := 1; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println("i =", i)
	}
}

func func_goto() {

	//等于java的标志
	for i := 1; i < 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i =", i)
	}
END:
	fmt.Println("end")
}

/*
	默认每个 case 带有 break
	case 中可以有多个选项
	fallthrough 等于没有break
*/
func func_switch() {
	i := 3
	fmt.Printf("当 i = %d 时\n", i)

	switch i {
	case 1:
		fmt.Println("输出i =", 1)
	case 2:
		fmt.Println("输出i =", 2)
	case 3:
		fmt.Println("输出i =", 3)
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出i =", "4 or 5 or 6")
	default:
		fmt.Println("输出 i=", "xxx")
	}
}
