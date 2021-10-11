package main

import "fmt"

func main(){
	// declare variables
	var a, b int

	// assign variables
	a, b = 5, 10

	// arithmetic operation
	c:=a+b
	fmt.Printf("%v + %v = %v\n", a, b, c)

	// subtraction
	d:=a-b
	fmt.Printf("%v - %v = %v\n", a, b, d)

	// division
	e:=float32(a) / float32(b)
	fmt.Printf("%v / %v = %v\n", a, b, e)

	// multiplication
	f:=a*b
	fmt.Printf("%v * %v = %v\n", a, b, f)
}