package main

import "fmt"

func main() {
	countryGDP := map[string]int{"Luxeburgo":131, "Switzerland":94, "Sierra Leone": 0, "Mexico":9}
	var ok bool
	_, ok = countryGDP["Sierra Leone"]
	if ok {
		fmt.Println("GDP assigned for Sierra Leone")
	} else {
		fmt.Println("GDP not assigned for Sierra Leone")
	}
	_, ok = countryGDP["Germany"]
	if ok {
		fmt.Println("GDP assigned for Germany")
	} else {
		fmt.Println("GDP not assigned for Germany")
	}
}
