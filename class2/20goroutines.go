package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 100; i++{
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 100; i++{
		fmt.Print("b")
	}
}

func c() {
	for i := 0; i < 100; i++{
		fmt.Print("c")
	}
}
/*
func main() {
   a()
   b()
   c()
   fmt.Println("End main() function")
}
*/

func main() {
	go a()
	go b()
	go c()
	time.Sleep(time.Second)
	fmt.Println("End main() function")
}
