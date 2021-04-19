package main

import "fmt"

func main() {
	var myMap map[string] int
	fmt.Printf("%#v\n", myMap)
	var myMap2 map[string] int = make(map[string]int)
	fmt.Printf("%#v\n", myMap2)
	myMap["I"] = 1
	fmt.Printf("%#v\n", myMap)
}
