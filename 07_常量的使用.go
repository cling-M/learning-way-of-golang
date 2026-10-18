package main

import "fmt"

func main(){
	//变量：程序运行期间 可以改变的量 变量声明需要 var 关键字
	//常量：程序运行期间 不可以改变的量 常量声明需要 const 关键字

	const a int = 10
	//a = 20  //err 常量不允许修改
	fmt.Println(a)

	const b = 10.1 //常量的定义不允许使用冒号等号
	fmt.Printf("b type is %T\n",b)
	fmt.Println(b)
}