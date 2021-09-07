package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds :=time.Now().Unix() // Number of seconds since January 1, 1970
	//fmt.Println(reflect.TypeOf(seconds))
	rand.Seed(seconds)
	myNumber := rand.Intn(3)-1
	fmt.Println(myNumber)
	if myNumber < 0 {
		fmt.Println("myNumber is negative")
	} else if myNumber == 0{
		fmt.Println("my Number is equal to cero")
	} else {
		fmt.Println("myNumber is positive")
	}
}
