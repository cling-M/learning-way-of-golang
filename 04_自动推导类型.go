package main

import "fmt"

func main(){
	//赋值之前必须提前声明变量 不然会出现undefined错误
	var a int
	a =10
	fmt.Println("a =",a)

	//:=会先给变量开始声明类型 再给变量进行赋值
	b := 20
	fmt.Println("b =",b)
	//前面有一个变量b 那么就不能再生成一个变量b 也就是说不能再直接写一个 b:= 60
	//但是赋值是可以的 就是不可以让他重新做一个自动推导
	b = 30
	fmt.Println("b =",b)
}