package main

import "fmt" //导入包后必须要使用

func main(){
	//变量就是 程序当中可以改变的量
	//声明格式 var 变量名 类型，变量声明了 必须要使用
	var a int
	//只是声明 但是没有初始化的变量默认是0
	//同一个{}内 声明的变量是唯一的 不可以重复定义
	fmt.Println("a =",a)
	//可以同时声明多个变量
	var b,c,d int
	fmt.Println(b,c,d)
	
	//变量赋值
	a =10
	fmt.Println("a =",a)
	
	//变量初始化：声明变量的时候同时赋值 
	var e int = 10 //初始化
	e = 20 //赋值
	fmt.Println("e =",e)

	//自动推导类型 必须初始化 根据初始化的数值进行推导类型（常用）
	f := 30
	//%T的意思就是 打印变量的类型
	fmt.Printf("c type is %T\n",f)
}