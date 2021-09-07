package main

import (
	"fmt"
	"sync"
)

var doOnce sync.Once


func main() {
	DoSomething()
	DoSomething()
	//fmt.Println(reflect.TypeOf(doOnce))
}

func DoSomething() {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
	})
	fmt.Println("Run this every time")
}
