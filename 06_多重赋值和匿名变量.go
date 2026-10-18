package main

import "fmt"

func test()(a,b,c int){
	return 1,2,3
}

func main() {
	//a := 10
	//b := 20
	//c := 30
	a,b := 10,20
	
	//交换两个变量的值
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a = %d,b=%d\n",a,b)

	i,j := 30,40
	i,j = j,i
	fmt.Printf("i = %d,j=%d\n",i,j)

	i = 50 
	j = 60
	//匿名变量 给i的值传送给temp j的值自动丢弃，只有配合函数的返回值的时候 才有优势
	temp,_ = i,j
	fmt.Printf("temp = %d\n",temp)

	var d,e,f int
	d,e,f = test()    //返回的是1，2，3 
	fmt.Println(d,e,f)
	//匿名变量
	_ , e , _ = test()
	fmt.Println(e)
}