package main

import "fmt"
func main() {
	type myStructType struct{
		figure string
		area float32
	}
	var myStruct myStructType
	myStruct.figure = "Circle"
	myStruct.area = 5
	fmt.Println(myStruct)
	var myStruct2 myStructType
	myStruct2.figure = "Rectangle"
	myStruct2.area = 10
	fmt.Println(myStruct2)
	myVar :=myStructType{"Square", 10.2}
	fmt.Println(myVar.figure)
	fmt.Println(myVar.area)
}

