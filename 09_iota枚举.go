package main

import "fmt"

func main(){
	//1，iota是常量自动生成器 每隔一行 iota就会加一
	//2，iota给常量赋值使用
	const(
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d,b = %d,c = %d\n",a,b,c)
	
	//3，iota遇到const会重置为0
	const d = iota
	fmt.Printf("d = %d\n",d)
	
	//4,可以只写一个iota
	const(
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1=%d,b1=%d,c1=%d\n",a,b,c)

	//5，同一行的值是一样的
	const(
		i = iota
		j1,j2,j3 = iota,iota,iota
		k =iota
	)
	fmt.Println(i,j1,j2,j3,k)
	} 