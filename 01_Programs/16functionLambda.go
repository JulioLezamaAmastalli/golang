package main

import (
	"fmt"
	"reflect"
)

func main() {
	func(){
	fmt.Println("Hello from a Lambda Function!")
}()

value := func() {
	fmt.Println("Hello, Go!")
}
value()
fmt.Println(reflect.TypeOf(value))
}
