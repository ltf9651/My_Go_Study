package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	default:
		return 0, fmt.Errorf("Wrong:%s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funcName:%s, args:%d, %d \n", opName, a, b)
	return op(a, b)
}

func div(a, b int) (r1, r2 int) {
	return a / b, a % b
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) {
	b, a = a, b
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(5, 3, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q1, _ := div(13, 3)
	fmt.Println(q1)
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3))
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
	a, b = swap3(a, b)
	fmt.Println(a, b)
}
