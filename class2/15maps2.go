package main

import "fmt"

func main() {
	countryGDP := map[string]int{"Luxeburgo":131, "Switzerland":94, "Sierra Leone": 0, "Mexico":9}
	fmt.Println("Sierra Leone", countryGDP["Sierra Leone"])
	fmt.Println("Germany", countryGDP["Germany"])
	var value int
	var ok bool
	value, ok = countryGDP["Sierra Leone"]
	fmt.Println("Sierra Leone",value, ok)
	value, ok = countryGDP["Germany"]
	fmt.Println("Germany",value, ok)
}
