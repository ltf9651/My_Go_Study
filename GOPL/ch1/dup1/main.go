//查找重复行
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		// 每次读取一行，该行当作map，其对应值递增，等价于
		//	line := line.Text()
		//	counts[line] = counts[line] ++
		// Text重复时对应的line + 1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
			// line 行重复 n 次
		}
	}
}

//%d 			十进制整数
// %x, %o, %b 	十六进制，八进制，二进制整数。
// %f, %g, %e 	浮点数： 3.141593 3.141592653589
// %t 			布尔：true或false
// %c 			字符（rune） (Unicode码点)
// %s 			字符串
// %q 			带双引号的字符串"abc"或带单引号的字符'a'
// %v 			变量的自然形式（natural format）
// %T 			变量的类型
// %% 			字面上的百分号标志（无操作
//指针是可见的内存地址，&操作符可以返回一个变量的内存地址，并且*操作符可以获取指针指向的变量内容
