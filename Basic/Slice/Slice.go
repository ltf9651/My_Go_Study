package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func slice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr[2:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])

	updateSlice(arr[2:])
	fmt.Println(arr[:])

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len=%d, cap=%d \n", s2, len(s2), cap(s2))

	s3 := append(s2, 11)
	fmt.Println(s3)
	fmt.Println(arr[:])
}

func opt() {
	var s []int // s == nil
	for index := 0; index < 100; index++ {
		printSlice(s)
		s = append(s, 2*index+1)
	}
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))
}

func main() {
	opt()
	s := make([]int, 10, 32)
	fmt.Println(s)
	printSlice(s)

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 20)
	copy(s2, arr)
	fmt.Println(s2)
	printSlice(s2)

	s2 = append(s2[:3], s2[4:]...) //delete 3
	fmt.Println(s2)

	front := s2[0]
	s2 = s2[1:]
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail, s2)
}
