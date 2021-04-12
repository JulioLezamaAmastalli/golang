package main

import "fmt"

func main() {
	//var myMap map[string] int
	//fmt.Printf("%#v\n", myMap)
	//myMap["I"] = 1
	var myMap map[string] int = make(map[string]int)
	fmt.Printf("%#v\n", myMap)
	myMap["I"] = 1
	fmt.Printf("%#v\n", myMap)
}
