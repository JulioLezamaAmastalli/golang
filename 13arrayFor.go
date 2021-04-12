package main

import "fmt"

func main()  {
	languages := [6]string{"Go", "Ruby", "Cobol", "Ada", "Swift", "Scala"}
	for index, value := range languages {
		fmt.Println(index, value)
	}
	// 7,5
	// blank identifier
}
