package main

import "fmt"

func main(){
	//1,声明变量
	var a bool
	a = true
	fmt.Println("a = ",a)

	//2，自动推导类型	
	b := false
	fmt.Println("b =",b)

	var c = true
	fmt.Println("c =",c)

	//3,没有赋值的时候 初始值是false
	var e bool
	fmt.Println("e =",e)
}