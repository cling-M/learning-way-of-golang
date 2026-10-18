package main

import "fmt"

func main(){
	var ch byte
	var str string
	
	//字符
	//1，单引号
	//2，字符往往都只有一个字符 ， 转义字符除外
	ch = 'a'
	fmt.Println("ch = ",ch)

	//字符串
	//1，双引号
	//2，由一个或者是多个字符组成的
	//3，结尾有一个转义字符'\0'
	str = "a" //它有两个字符 一个是 'a' 还有一个是'\0'
	fmt.Println("str =",str)

	str2 := "012345"
	fmt.Printf("str2[0] = %c,str2[1] = %c\n",str2[0],str2[1])
}