package main

import "fmt"

func main(){
	var ch byte //声明字符类型
	ch = 65

	//格式化输出 %c以字符方式打印 %d以整形方式打印
	fmt.Printf("%c,%d\n",ch,ch)

	ch = 'A'    //字符用单引号的
	fmt.Printf("%c,%d\n",ch,ch)

	//大写转小写
	fmt.Printf("大写：%d 小写：%d\n",'A','a')
	ch = ch +32
	fmt.Printf("%c,%d\n",ch,ch)
}