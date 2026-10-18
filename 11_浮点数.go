package main

import "fmt"

func main(){
	//声明变量
	var a float32
	a = 3.14
	fmt.Println("a = ",a)

	//自动推导类型 默认是float64
	b := 5.20
	fmt.Printf("b type is %T\n",b)

	//float64比float32计算出来更加的准确 
}