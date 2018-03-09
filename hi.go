package main

//#cgo CFLAGS: -I/Users/wfps/buru/cups/cups
//#include<stdlib.h>
//#include "config.h"
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func main() {
	fmt.Println(Random())
}
