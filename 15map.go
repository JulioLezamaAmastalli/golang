package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["Ruby"] = 2
	myMap["Ada"] = 1
	myMap["Cobol"] = 3
	myMap["Go"] = 4
	fmt.Println(myMap)
	rank := map[string]int{"First":1, "Second":2, "Third":3}
	fmt.Println(rank)
}