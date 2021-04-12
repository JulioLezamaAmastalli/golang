package main

import "fmt"

func main() {
	var a int32
	var b float32
	a = 10
	b = 10
	a = b + a // compiler error
	fmt.Println(a)
	//var d int32 = 0
	//fmt.Println(a / d)
}
