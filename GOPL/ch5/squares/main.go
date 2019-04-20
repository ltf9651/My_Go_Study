package main

import "fmt"

//返回下一个属的平方根
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	//匿名函数中存在变量引用，记录了局部变量X的状态
}
