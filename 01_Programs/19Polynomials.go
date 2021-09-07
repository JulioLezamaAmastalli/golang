package main

import (
	"fmt"
	"polynomial"
)

func main() {
	p := polynomial.ToPol(polynomial.Term{4, 3}, polynomial.Term{3, 2}, polynomial.Term{2, 1}, polynomial.Term{1, 0})
	q := polynomial.ToPol(polynomial.Term{3, 2}, polynomial.Term{5, 0})
	z := polynomial.ToPol(polynomial.Term{0, 0})
	fmt.Printf("%-15s %s %s\n", "zero(x)","=", polynomial.ToString(z))
	fmt.Printf("%-15s %s %s\n", "p(x)","=", polynomial.ToString(p))
	fmt.Printf("%-15s %s %s\n", "q(x)", "=", polynomial.ToString(q))
	fmt.Printf("%-15s %s %s\n", "p(x) + q(x)", "=", polynomial.ToString(polynomial.Plus(p, q)))
	fmt.Printf("%-15s %s %s\n", "p(x) * q(x)", "=", polynomial.ToString(polynomial.Times(p, q)))
	fmt.Printf("%-15s %s %s\n", "p(q(x))", "=", polynomial.ToString(polynomial.Compose(p,q)))
	fmt.Printf("%-15s %s %s\n", "0 - p(x)", "=", polynomial.ToString(polynomial.Minus(z, p)))
	fmt.Printf("%-15s %s %d\n", "p(3)", "=", polynomial.Evaluate(p,3))
	fmt.Printf("%-15s %s %s\n", "p'(x)", "=", polynomial.ToString(polynomial.Differentiate(p)))
	fmt.Printf("%-15s %s %s\n", "p''(x)", "=", polynomial.ToString(polynomial.Differentiate(polynomial.Differentiate(p))))
}
