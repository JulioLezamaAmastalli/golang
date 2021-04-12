package main

import "fmt"

func main() {
	var mySlice []int
	fmt.Println(mySlice)
	var mySlice2 []int
	mySlice2 = make([]int, 4)
	fmt.Println(mySlice2)
	var mySlice3 = []string{"a","b","c"}
	fmt.Println(mySlice3)
	mySlice4 := []string{"I","II","III"}
	fmt.Println(mySlice4)
	array1 :=[]string{"a","b","c","d","e"}
	mySlice5 := array1[2:4]
	fmt.Println(mySlice5)
}


