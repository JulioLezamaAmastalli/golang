package main

import (
	"fmt"
	"math/cmplx"
)

func main()  {
	var myComplexNumber complex64 = 1 + 2i
	fmt.Println(myComplexNumber)
	fmt.Println(real(myComplexNumber))
	fmt.Println(imag(myComplexNumber))
	fmt.Println(cmplx.Abs(3+4i))
}

