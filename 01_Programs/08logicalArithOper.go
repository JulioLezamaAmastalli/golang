package main

import "fmt"

func main() {
	var a  = 20
	var b  = 10
	var c  = 15
	var d  = 5
	var e int

	e = a + b * c / d
	fmt.Println("Value of a + b * c / d is :",  e )

	e = (a + b * c) / d
	fmt.Println("Value of (a + b * c) / d is  :" ,  e )

	e = ((a + b) * c)  / d
	fmt.Println("Value of ((a + b) * c) / d) is  :",  e )
}
