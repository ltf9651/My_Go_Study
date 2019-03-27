package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*
$ go build GOPL/ch2/echo4
$ ./echo4 a bc defa bc def
$ ./echo4 -s / a bc defa/bc/def
$ ./echo4 -n a bc defa bc def$
$ ./echo4 -help
	Usage of ./echo4:
		-n omit trailing newline
		-s stringseparator (default " ")
*/
