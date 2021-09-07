package main

import (
	"fmt"
)

func greeting(myChannel chan string)  {
	myChannel <- "Hi!"
}

func main() {
	//var myChannel chan string
	//myChannel = make(chan string)
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)
	//value := <-myChannel
	//fmt.Println(value)
}
