package main

import "fmt"

func main(){
	var str string
	str = "I love you"
	fmt.Println("str is :",str)

	//自动推导类型
	str1 := "cling"
	fmt.Println("My name is :",str1)
	fmt.Printf("str1`s type is:%T\n",str1)

	//内建函数 len()来检测这个字符串当中有多少字符
	str2 := "12345678"
	fmt.Println("len(str2)=",len(str2))
}