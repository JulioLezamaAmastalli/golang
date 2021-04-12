package main

import "fmt"

func main() {
	var a  = 20
	var b  = 10
	var c  = 15
	var d  = 5
	var e int
	// Formatting verbs
	// verb
	// value
	// Formatting value widths
	e = a + b * c / d
	fmt.Printf("%40s %d\n", "Value of a + b * c / d is :",  e )

	e = (a + b * c) / d
	fmt.Printf("%40s %d\n", "Value of (a + b * c) / d is :" ,  e )

	e = ((a + b) * c)  / d
	fmt.Printf("%40s %d\n", "Value of ((a + b) * c) / d) is :",  e )
}
