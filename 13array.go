package main

import "fmt"

func main() {
	var myArray [5]int
	fmt.Println("myArray:", myArray)
	myArray[4] = 10
	fmt.Println("myArray initialized on index 4:", myArray)
	fmt.Println("The length of myArray:", len(myArray))
	myArray2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("myArray2 declared and initialized:", myArray2)
	var romanNumbers [3]string = [3]string{"I", "II", "III"}
	// romanNumbers := [3]string{"I", "II", "III"}
	fmt.Println("romanNumbers:", romanNumbers)
}