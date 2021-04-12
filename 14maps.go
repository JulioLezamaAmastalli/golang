package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = make(map[string]int) // myMap := make(map[string]float64)
	myMap["Ruby"] = 2
	myMap["Adda"] = 1
	myMap["Cobol"] = 3
	myMap["Go"] = 4
	fmt.Println(myMap)
	ranks := map[string]int{"First":1, "Second":2, "Third":3}
	fmt.Println(ranks)
}
