package main

import "fmt"

func main()  {
	var a int
	fmt.Printf("请你输入a的值\n")
	// fmt.Scanf("%d",&a)  //一定要记得给a取地址
	fmt.Scan(&a) 
	fmt.Println("a = ",a)
}