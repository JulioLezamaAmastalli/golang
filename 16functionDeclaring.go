package main

import "fmt"

func plusInt(a int, b int) int {
	return a + b
}

func main() {
	res := plusInt(1, 2)
	fmt.Println("1+2 =", res)
}
// function here