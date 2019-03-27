// defer 函数结束时发生  open/close lock/unlock PrintHeader/PrintFooter
package main

import (
	"My_Go_Study/Basic/ErrorHandle/fib"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	err := errors.New("this is an error")
	panic(err)
}

func tryDefer2() {
	for index := 0; index < 20; index++ {
		defer fmt.Println(index)
		if index == 15 {
			// panic("too many")
			fmt.Println("error")
			return
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		// panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file) //先写入内存
	defer writer.Flush()            //写入文件

	f := fib.Fibonacci()
	for index := 0; index < 20; index++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
	tryDefer2()
}
