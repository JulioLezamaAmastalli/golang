package main

import "fmt"
func main() {
	var myStruct struct{
		figure string
		area float32
	}
	fmt.Println(myStruct)
	myStruct.figure = "Circle"
	myStruct.area = 5
	fmt.Println(myStruct)
}
