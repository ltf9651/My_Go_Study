package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 8}
	reverse(a[:])
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//测试一个slice是否是空使用len(s) == 0来判断，而不应该用s == nil来判断。
