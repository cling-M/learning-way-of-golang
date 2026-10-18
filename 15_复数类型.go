package main

import "fmt"

func main(){
	var t complex128 //声明
	t = 2.1 + 3.14i  //赋值
	fmt.Println("t = ",t)

	//自动推导类型
	t2 := 5 + 0.2i
	fmt.Println("t2 =",t2)
	fmt.Printf("t2`type is %T\n",t2)

	//通过内建函数 来取得实部虚部
	fmt.Println("real(t2) is :",real(t2),",imag(t2) is :",imag(t2),"\n")
}