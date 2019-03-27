package main

import (
	"fmt"
	"math"
)

//包内变量
var (
	p1 = 1
	p2 = "kkk"
	p3 = false
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a, c int = 3, 4
	var b = "abc"
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	var a, c, b = 3, true, "abc"
	fmt.Println(a, b, c)
}

func variableShorter() {
	a, c, b := 3, true, "abc"
	b = "qw"
	fmt.Println(a, b, c)
}

func triangle() {
	var a, b int = 3, 4
	// var c int
	//c = math.Sqrt(a*a + b*b) Sqrt需要传float
	//c = math.Sqrt(float64(a*a + b*b)) Sqrt返回float, c必为int
	// c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(calTriangle(a, b))
}

func calTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4 // 常量的数值可以作为各数据类型使用
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enum() {
	// const (
	// 	php    = 0
	// 	python = 1
	// 	golang = 2
	// )
	const (
		php = iota
		_
		python
		golang
		java
	)

	const (
		b = 1 << (iota)
		b1
		b2
		b3
		b4
		b5
	)

	fmt.Println(php, java, python, golang)
	fmt.Println(b, b1, b2, b3, b4, b5)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	consts()
	enum()
}
