package main

import "fmt"

func main() {
	var num int

	//支持一个初始化语句 初始化语句和变量本身用分号分隔开
	switch num = 1; num {
	case 1:
		fmt.Println("cling")
		//不跳出switch语句 后面的无条件执行
	case 2, 3:
		fmt.Println("slang")
		fallthrough
	case 4:
		fmt.Println("cling and slang")
		fallthrough //go语言保留了break关键字，来跳出switch语言，不写默认包含
	default:
		fmt.Println("cling love slang")

	}

	score := 85
	switch { //switch可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("slang 比 cling 爱对方多一点")
	case score == 85:
		fmt.Println("cling 和 slang 说：爱是相互的")
	case score < 85:
		fmt.Println("cling 比 slang 爱对方多一点")
	default:
		fmt.Println("cling 和 slang 很恩爱呢")
	}
}
