package main

import "fmt"

func main() {
	func(message string){
		fmt.Println(message)
	}("Hello from a Lambda Function")

	for i := 0; i < 4; i++ {
		func(i int) { fmt.Println(i) }(i)
	}
}