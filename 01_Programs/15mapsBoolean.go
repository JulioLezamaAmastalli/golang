package main

import "fmt"

func main() {
	romanNumbers := map[string]int{"I":1, "II":2, "III":3}
	var value int
	var ok bool
	value, ok = romanNumbers["I"]
	fmt.Println(value, ok)
	value, ok = romanNumbers["IV"]
	fmt.Println(value, ok)
	fmt.Println(romanNumbers)
	delete(romanNumbers, "II")
	fmt.Println(romanNumbers)
}
