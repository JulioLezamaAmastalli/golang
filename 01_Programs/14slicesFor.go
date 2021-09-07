package main

import "fmt"

func main() {
	romanNumbers := []string{"I", "II", "III"}
	for i := 0; i <len(romanNumbers); i++{
		fmt.Println(romanNumbers[i])
	}
	for key, value := range romanNumbers{
		fmt.Println(key, value)
		// range and blank identifier
	}
}
