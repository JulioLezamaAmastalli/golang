package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var day string // Declaration
	day = "Monday"
	fmt.Println("Today is ", day)
	var length, width int // Multiple declaration
	length, width = 10, 20
	fmt.Println(length, width)
	var letter = "A" // Declaration and assignment together
	fmt.Println(letter)
	food := "Cake" // Short variable declaration
	fmt.Println(food)
	a, b, c := "A", 1, false
	fmt.Println(a,b,c)
	fmt.Println(reflect.TypeOf(b))
	_x := 9
	//var x int
	fmt.Println(_x) //without assigning a value
	const myConst  = 8
	fmt.Println(myConst)
}
