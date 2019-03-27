package main

import (
	"My_Go_Study/GOPL/ch2/tempconv"
	"fmt"
	"reflect"
)

type Celsius float64     //摄氏度
type Fahrenherit float64 //华氏度

const (
	AbosoluteZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	BoilingC       Celsius = 100
)

func CToF(c Celsius) Fahrenherit {
	return Fahrenherit(c*9/5 + 32)
}

func FToC(f Fahrenherit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var c Celsius
	var f Fahrenherit
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(f))
	//fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // true
	fmt.Println(tempconv.CToF(tempconv.FreezingC))
}
