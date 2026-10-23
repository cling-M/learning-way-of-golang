package main

import "fmt"

func main() {
	s := "slang"
	if s == "slang" {
		fmt.Println("slang和cling")
	}
	// if支持一个初始化语句，初始化语句和判断条件用分号分隔
	if a := 10; a == 10 {
		fmt.Println("cling和slang")
	}
}
