package main

import "fmt"

func main() {
	//1
	a := 11
	if a == 10 {
		fmt.Println("cling 和 slang")
	} else {
		fmt.Println("slang 和 cling")
	}

	// 2
	if a := 10; a == 10 {
		fmt.Println("cling love slang")
	} else {
		fmt.Println("slang love cling")
	}

	a = 10
	if a == 10 {
		fmt.Println("cling and slang")
	} else if a > 10 {
		fmt.Println("slang and cling")
	} else if a < 10 {
		fmt.Println("slang cling")
	} else {
		fmt.Println("cling slang")
	}
}
