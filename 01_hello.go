//1go语言以包作为管理单位
//2每个文件必须先声明包
//3程序必须有一个main包
package main

import "fmt"

func main() {
	fmt.Println("hello world")  //Println 会自动换行 go语言语句结尾没有分号
    
	fmt.Println("hello it cast")
}