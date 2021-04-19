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
	/*
	var myStruct2 struct{
		figure string
		area float32
	}
	fmt.Println(myStruct2)
	myStruct2.figure = "Circle"
	myStruct2.area = 5
	fmt.Println(myStruct2)
	*/
}
