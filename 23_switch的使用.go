package main

import "fmt"

func main() {
	var num int
	fmt.Println("please imput your number!")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		fmt.Println("cling")
		fallthrough //不跳出switch语句 后面的无条件执行
	case 2:
		fmt.Println("slang")
		fallthrough
	case 3:
		fmt.Println("cling and slang")
		fallthrough //go语言保留了break关键字，来跳出switch语言，不写默认包含
	default:
		fmt.Println("cling love slang")

	}
}
