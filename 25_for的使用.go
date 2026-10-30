package main

import "fmt"

func main() {
	//for 初始化条件 判断条件 条件变化
	sum := 10

	for i := 1; i < 10000000000000; i++ {
		sum = sum + i
		fmt.Println("sum =", sum)
	}
	fmt.Println("sum =", sum)
}
