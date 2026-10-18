package main

import "fmt"

func main(){
	a := 10
	b := 520.1314
	c := 'c'
	d := "I love you"

	fmt.Printf("a:%T,b:%T,c:%T,d:%T\n",a,b,c,d)

	fmt.Printf("a:%d,b:%f,c:%c,d:%s\n",a,b,c,d)
	//用%v来自动匹配格式
	fmt.Printf("a:%v,b:%v,c:%v,d:%v\n",a,b,c,d)
}