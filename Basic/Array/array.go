package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray1(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 5, 7}
	arr3 := [...]int{2, 5, 3, 3, 5}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	// Go语言更多用Slice代替Array和指针
	printArray1(arr1)
	fmt.Println(arr1)
	printArray(&arr1)
	fmt.Println(arr1)
}
