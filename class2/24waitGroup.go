package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Wait for a collection of goroutines to run

func sayHi(names []string) {
	defer wg.Done() // Decrease the number of goroutines to wait by one
	for _, name := range names {
		fmt.Printf("Hi %s\n\r", name)
		time.Sleep(1 * time.Second)
	}
}

func sayBye(names []string) {
	defer wg.Done() // Decrease the number of goroutines to wait by one
	for _, name := range names {
		fmt.Printf("Bye %s\n\r", name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg.Add(2) // Add the number of goroutines to wait

	nombres := []string{"Andres", "Brandon", "Luis",
		"Braulio", "Salvador", "Diego"}

	go sayHi(nombres)
	go sayBye(nombres)

	wg.Wait() // Wait until the goroutines run
	fmt.Println("End")
}
