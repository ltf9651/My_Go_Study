package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 11111@qq.com@qq.com
email2 is 2qsd@126.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}
