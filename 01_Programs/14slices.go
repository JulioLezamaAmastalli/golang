package main

import "fmt"

func main() {
	var mySlice []int // declaring
	fmt.Printf("mySlice: %#v\n",mySlice)
	var mySlice2 []int // creating
	mySlice2 = make([]int, 4)
	fmt.Printf("mySlice2: %#v\n", mySlice2)
	var mySlice3 = []string{"a","b","c"}
	fmt.Printf("mySlice3: %#v\n", mySlice3)
	mySlice4 := []string{"I","II","III"}
	fmt.Printf("mySlice4: %#v\n",mySlice4)
	array1 :=[5]string{"a","b","c","d","e"}
	mySlice5 := array1[2:4]
	fmt.Printf("mySlice5: %#v\n", mySlice5)
	//mySlice5[0] = "e"
	//fmt.Printf("mySlice5: %#v\n", mySlice5)
	//fmt.Printf("array1: %#v\n", array1)
	//array1[2] = "K"
	//fmt.Printf("mySlice5: %#v\n", mySlice5)
	//fmt.Printf("array1: %#v\n", array1)
}


