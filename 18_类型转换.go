package main

import "fmt"

func main() {
	// 这种不能转换的类型，叫做不兼容类型
	var flag bool
	flag = true

	fmt.Printf("flag : %t\n", flag)

	// bool 类型不可以转换成int类型
	// fmt.Printf("flag : %d\n", int(flag))

	// 0就是假，1就是真
	// 整形也不能转换成bool类型
	// flag = 1

	var ch byte
	ch = 'a' //字符类型本质上就是整形
	var t int
	t = int(ch) //把ch的值取出来 然后转换成int类型 然后再赋值给t
	fmt.Printf("t:%d\n", t)
}
