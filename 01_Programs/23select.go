package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // in goroutine_select2.go
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump3(ch1)
	go pump2(ch2)
	go suck1(ch1, ch2)
	time.Sleep(1e9)
}
func pump3(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i 
	}
}
func suck1(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
