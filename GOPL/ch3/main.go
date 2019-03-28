package main

import "fmt"

func main() {
	var f float32 = 16777216
	var g float64 = 16777216
	fmt.Println(f == f+1, g == g+1)

	var f2 float64 = 212
	fmt.Println((f2 - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f2 - 32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f2 - 32)) // "100"; 5.0/9.0 is an untyped flo
}
