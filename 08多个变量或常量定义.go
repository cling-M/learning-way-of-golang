package main

import "fmt"

func main(){
	// 不同类型的定义
	// var a int = 10
	// var b float64 = 3.14
	var(
		a int
		b float64
    )
	a,b = 10,3.14
	fmt.Printf("a = %d\n",a)
	fmt.Printf("b = %f\n",b)

	//自动推导类型
    var(
		m =10
		y = 52
	)

	fmt.Println(m)
	fmt.Println(y)
	// const i int =10
	// const j float64 = 3.14
	const (
		i int = 10
		j float64 = 3.14
	)
	fmt.Println(i)
	fmt.Println(j)

	//const自动推导类型
	const(
		w = 13
		s = 14
	)
	fmt.Println(w)
	fmt.Println(s)
}